package models

import (
	"crypto/md5"
	"fmt"
	"strings"
)

/*
参考数据源：
    {
        "jsxm":"张三", //教师姓名
        "jsmc":"教学楼101", //教室名称
        "jssj":"10:00", //结束时间
        "kssj":"08:00", //开始时间
        "kkzc":"1", //开课周次，有三种已知格式1)a-b、2)a,b,c、3)a-b,c-d
        "kcsj":"10506", //课程时间，格式x0a0b，意为星期x的第a,b节上课
        "kcmc":"大学英语", //课程名称
        "sjbz":"0" //具体意义未知，据观察值为1时本课单周上，2时双周上
    }
*/
type Course struct {
	Cid      string // 该系统内部课程识别码，Name,Teacher,Location,Begin,End,Week 的 md5 和
	Name     string // 课程名称
	Teacher  string // 任课老师
	Location string // 上课地点
	Begin    string // 上课时间
	End      string // 下课时间
	Week     string // 开课周次
}

// NewCourse 返回给定 name,teacher,location,begin,end,week 所决定的 Course，cid 会在此完成计算
func NewCourse(name string, teacher string, location string, begin string, end string, week string) *Course {
	course := &Course{Name: name, Teacher: teacher, Location: location, Begin: begin, End: end, Week: week}

	// 计算 cid，Name,Teacher,Location,Begin,End,Week 的 md5 和
	sl := []string{name, teacher, location, begin, end, week}
	data := []byte(strings.Join(sl, ""))
	cid := fmt.Sprintf("%x", md5.Sum(data))
	course.Cid = cid

	return course
}

func testNewCourse() { // TODO: delete this func
	c := NewCourse("上帝学", "耶稣", "巴黎圣母院2楼小室", "08:00", "09:50", "1-12")
	fmt.Println(c, c.Cid)
}
