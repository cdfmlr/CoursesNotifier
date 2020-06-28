/*
 * Copyright 2020 CDFMLR
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package qzclient

import (
	"example.com/CoursesNotifier/data"
	"example.com/CoursesNotifier/models"
	"github.com/cdfmlr/qzgo"
	"log"
	"math/rand"
	"strconv"
	"time"
)

const SCHOOL = "ncepu"

type Client struct {
	qzgo.Client
	Student       models.Student
	CurrentXnxqId string
	CurrentWeek   string
	Courses       map[string]models.Course
}

func New(student models.Student) *Client {
	if student.Sid == "" {
		log.Fatal("student.Sid should not be Empty!")
	}

	client := &Client{Student: student}
	client.School = SCHOOL
	client.Xh = client.Student.Sid
	client.Pwd = client.Student.Pwd

	return client
}

// Login 登录强智系统，获取操作 token，在该 token 过期之前可以做其他操作
func (c *Client) Login() (authUserRespBody *qzgo.AuthUserRespBody, err error) {
	authUserRespBody, err = c.AuthUser()
	if err != nil {
		log.Println(err)
	}
	return authUserRespBody, err
}

// FetchCurrentTime 获取当前学期、周次
func (c *Client) FetchCurrentTime() error {
	current := time.Now().Format("2006-01-02")
	getCurrentTimeRespBody, err := c.GetCurrentTime(current)
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
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	getKbcxAzcRespBodyItems, err := c.GetKbcxAzc(c.Student.Sid, c.CurrentXnxqId, strconv.Itoa(week))
	if err != nil {
		log.Println(err)
	}
	courses := make([]models.Course, 0)
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

	// TODO: 在 database 里实现一个 insertMany，省得这里麻烦，且重复连接数据库
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

	// TODO: 在 database 里实现一个 insertMany，省得这里麻烦，且重复连接数据库
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
