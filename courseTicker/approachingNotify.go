package courseTicker

import (
	"example.com/CoursesNotifier/data"
	"example.com/CoursesNotifier/models"
	"example.com/CoursesNotifier/qz/qzapi"
	"fmt"
	"log"
	"math"
	"regexp"
	"strings"
	"time"
)

const secondToWeek = 60 * 60 * 24 * 7

// NotifyApproachingCourses 通知快要开始上的课。
//
// 具体来说，该函数从数据库 databaseSource 中，
// 找出 minuteBeforeCourseToNotify 分钟内要开始的课，
// 以及上这些课的学生，调用 notifiers 进行通知。
func (ct *CoursesTicker) NotifyApproachingCourses() {
	// 今天星期几
	currentWeek := getCurrentWeek(ct.databaseSource)
	// 今天第几周
	todayQzWeekday := qzapi.TimeWeekToQzWeekday(time.Now().Weekday())

	// 最近一个可能上课的时间
	nearCourseTime := getNearestBeginTime(ct.databaseSource)
	nearCourseTimeStr := nearCourseTime.Format("15:04")

	// 现在离上课还早着呢，就地返回，别去烦人家了
	if nearCourseTime.Sub(time.Now()).Minutes() > ct.minuteBeforeCourseToNotify {
		log.Printf(
			"Don't Notify: nearCourseTime(%s) - now(%s) > %v Minutes; \n",
			nearCourseTime,
			time.Now(),
			ct.minuteBeforeCourseToNotify,
		)
		return
	}

	// 这学期星期几、几点开始上的所有课
	cdb := data.NewCourseDatabase(ct.databaseSource)
	coursesApproaching, _ := cdb.GetCoursesOnTime(int(todayQzWeekday), nearCourseTimeStr)

	// 细化到*本周*星期几、几点开始上的课
	coursesApproachingInWeek := make([]models.Course, 0)
	for _, c := range coursesApproaching {
		if ok, _ := isCourseInWeek(&c, currentWeek); ok {
			coursesApproachingInWeek = append(coursesApproachingInWeek, c)
		}
	}

	// 找上这些课的学生，通知ta们要开始上课了
	rdb := data.NewStudentCourseRelationshipDatabase(ct.databaseSource)
	sdb := data.NewStudentDatabase(ct.databaseSource)
	for _, c := range coursesApproachingInWeek {
		relations, _ := rdb.GetRelationshipsOfCourse(c.Cid)
		for _, r := range relations {
			s, _ := sdb.GetStudent(r.Sid)
			for _, n := range ct.notifiers {
				n.Notify(s, &c)
			}
		}
	}
}

// getNearestBeginTime 获取距离当前最近的课程时间，出错时返回当前时间的15分钟后
func getNearestBeginTime(databaseSource string) time.Time {
	loc, _ := time.LoadLocation("PRC") // 时区为 CST-8

	nearest, _ := time.ParseInLocation("Jan 2 2006", "Jan 2 2200", loc) // Duration 只能放 290 年，我不信 180 年后还有人用这个烂系统
	now := time.Now()

	courseBeginTimes, err := _getPossibleCourseBeginTime(databaseSource)

	for _, cbt := range courseBeginTimes {
		s := cbt.Sub(now)
		if s > 0 && s < nearest.Sub(now) {
			nearest = cbt
		}
	}
	if err != nil {
		log.Println(err)
		return time.Now().Add(15 * time.Minute)
	}
	return nearest
}

// _getPossibleCourseBeginTime 返回数据库中今、明两天的所有可能上课时间
func _getPossibleCourseBeginTime(databaseSource string) ([]time.Time, error) {
	loc, _ := time.LoadLocation("PRC") // 时区为 CST-8
	now := time.Now()

	cdb := data.NewCourseDatabase(databaseSource)
	courseBeginTimesStr, err := cdb.GetCoursesBeginTime()

	courseBeginTimes := make([]time.Time, 0)

	for _, cbtStr := range courseBeginTimesStr {
		cbtHM, _ := time.ParseInLocation("15:04", cbtStr, loc)
		cbtToday := cbtHM.AddDate(now.Year(), int(now.Month())-1, now.Day()-1)
		cbtTomorrow := cbtHM.AddDate(now.Year(), int(now.Month())-1, now.Day())
		courseBeginTimes = append(courseBeginTimes, cbtToday, cbtTomorrow)
	}

	if err != nil {
		log.Println(err)
		return []time.Time{}, err
	}

	return courseBeginTimes, err
}

// isCourseInWeek 判断一个 models.Course 是否在指定周次(week) 有课
func isCourseInWeek(course *models.Course, week int) (bool, error) {
	parts := strings.Split(course.Week, ",")

	var err error
	var match bool
	for _, w := range parts {
		if match, err = regexp.MatchString(`^(\d*?)-(\d*)$`, w); match {
			// a-b 型
			begin, end := 0, 0
			_, err = fmt.Sscanf(w, "%d-%d", &begin, &end)
			if week >= begin && week <= end {
				return true, nil
			}
		} else if match, err = regexp.MatchString(`^(\d*?)$`, w); match {
			// a 型
			begin := 0
			_, err = fmt.Sscanf(w, "%d", &begin)
			if week == begin {
				return true, nil
			}
		}
	}

	if err != nil {
		log.Println(err)
		return false, err
	}
	return false, nil
}

// getCurrentWeek 获取当前教学周次
func getCurrentWeek(databaseSource string) int {
	crtdb := data.NewCurrentDatabase(databaseSource)
	termBegin, err := crtdb.GetCurrentTermBeginDate()
	if err != nil {
		log.Println(err)
		return 0
	}
	diff := time.Since(termBegin)
	if diff < 0 {
		return 0
	}
	return 1 + _durationToWeek(diff)
}

// durationToWeek convert a duration into week
func _durationToWeek(duration time.Duration) int {
	return _roundTime(duration.Seconds() / secondToWeek)
}

// roundTime helps getting a reasonable int from a float, which is of great help when converting the duration into week
func _roundTime(input float64) int {
	var result float64

	if input < 0 {
		result = math.Ceil(input)
	} else {
		result = math.Floor(input)
	}

	// only interested in integer, ignore fractional
	i, _ := math.Modf(result)

	return int(i)
}
