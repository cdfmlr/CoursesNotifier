package qzclient

import (
	"example.com/CoursesNotifier/data"
	"example.com/CoursesNotifier/models"
	"example.com/CoursesNotifier/qzapi"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type Client struct {
	Student       models.Student
	token         string
	CurrentXnxqId string
	CurrentWeek   string
	Courses       map[string]models.Course
}

func NewClient(student models.Student) *Client {
	if student.Sid == "" {
		log.Fatal("student.Sid should not be Empty!")
	}
	return &Client{Student: student}
}

// AuthUser 登录强智系统，获取操作 token，在该 token 过期之前可以做其他操作
func (c *Client) AuthUser() (authUserRespBody *qzapi.AuthUserRespBody, err error) {
	authUserRespBody, err = qzapi.AuthUser(qzapi.SchoolNcepu, c.Student.Sid, c.Student.Pwd)
	if err != nil {
		log.Println(err)
	}
	c.token = authUserRespBody.Token
	return authUserRespBody, err
}

// FetchCurrentTime 获取当前学期、周次
func (c *Client) FetchCurrentTime() error {
	current := time.Now().Format("2006-01-02")
	getCurrentTimeRespBody, err := qzapi.GetCurrentTime(qzapi.SchoolNcepu, c.token, current)
	if err != nil {
		log.Println(err)
		return err
	}
	c.CurrentXnxqId = getCurrentTimeRespBody.Xnxqh
	c.CurrentWeek = strconv.Itoa(getCurrentTimeRespBody.Zc)
	return nil
}

// FetchAllTermCourses 获取整个学期的所有课程（要反爬虫，速度很慢）
func (c *Client) FetchAllTermCourses(ch chan []models.Course) error {
	const maxWeek = 21

	chweek := make(chan []models.Course, maxWeek+1)
	count := 0

	for zc := 1; zc <= maxWeek; zc++ { // 遍历所有周，得到完整学期课表
		c.FetchWeekCoursesSlowly(zc, chweek)
	}

LOOP:
	for {
		select {
		case getKbcxAzcRespBodyItems := <-chweek:
			c.appendCourse(getKbcxAzcRespBodyItems)
			count++
		default:
			if count >= maxWeek {
				break LOOP
			}
		}
	}

	var termCourseList []models.Course
	for _, v := range c.Courses {
		termCourseList = append(termCourseList, v)
	}
	if ch != nil {
		ch <- termCourseList
	}
	return nil
}

// FetchWeekCoursesSlowly 获取某一周的课程（要反爬虫，速度很慢）
func (c *Client) FetchWeekCoursesSlowly(week int, ch chan []models.Course) {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	getKbcxAzcRespBodyItems, err := qzapi.GetKbcxAzc(qzapi.SchoolNcepu, c.token, c.Student.Sid, c.CurrentXnxqId, strconv.Itoa(week))
	if err != nil {
		log.Println(err)
	}
	var courses []models.Course
	for _, v := range getKbcxAzcRespBodyItems {
		c := models.NewCourse(v.Kcmc, v.Jsxm, v.Jsmc, v.Kssj, v.Jssj, v.Kkzc, v.Kcsj)
		courses = append(courses, *c)
	}
	ch <- courses
}

// appendCourse appends Courses to Client without duplicates.
func (c *Client) appendCourse(courses []models.Course) {
	if c.Courses == nil {
		c.Courses = make(map[string]models.Course)
	}
	for _, v := range courses {
		c.Courses[v.Cid] = v
	}
}

// Save 将数据（student、course、s_c_relationship）保存到数据库
func (c *Client) Save(databaseSource string) (rowsAffected int64) {
	var totalAffected int64

	totalAffected += c.saveStudent(databaseSource)
	totalAffected += c.saveCourses(databaseSource)
	totalAffected += c.saveSCRelationships(databaseSource)

	return totalAffected
}

func (c *Client) saveStudent(databaseSource string) (rowsAffected int64) {
	sdb := data.NewStudentDatabase(databaseSource)
	affected, err := sdb.Insert(c.Student)
	if err != nil {
		log.Println(err)
	}
	return affected
}

func (c *Client) saveCourses(databaseSource string) (rowsAffected int64) {
	var totalAffected int64

	cdb := data.NewCourseDatabase(databaseSource)
	for _, v := range c.Courses {
		affected, err := cdb.Insert(v)
		if err != nil {
			log.Println(err)
		}
		totalAffected += affected
	}
	return totalAffected
}

func (c *Client) saveSCRelationships(databaseSource string) (rowsAffected int64) {
	var totalAffected int64

	relations := make([]models.Relationship, 0)
	for k := range c.Courses {
		relations = append(relations, *models.NewRelationship(c.Student.Sid, k))
	}

	rdb := data.NewStudentCourseRelationshipDatabase(databaseSource)
	for _, r := range relations {
		affected, err := rdb.Insert(r)
		if err != nil {
			log.Println(err)
		}
		totalAffected += affected
	}

	return totalAffected
}
