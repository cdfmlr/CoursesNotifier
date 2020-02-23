package models

import (
	"time"
)

type Student struct {
	Sid        string   // 学号
	Pwd        string   // 教务密码
	WxUser     string   // 微信用户公众号openid
	Courses    []Course // 课程列表
	CreateTime int64    // 创建时间
}

// NewStudent returns a Student of given sid and pwd
func NewStudent(sid string, pwd string, wxUser string) *Student {
	s := &Student{Sid: sid, Pwd: pwd, WxUser: wxUser}
	s.CreateTime = time.Now().Unix()
	return s
}

// IsLiving returns true when a Student is updated and living.
// This is measured via checking if the given expire is greater the time elapsed from the CreateTime of Student s.
// Giving a expire < 0 will always get true, means never expires.
func (s *Student) IsLiving(expire int64) bool {
	if expire < 0 {
		return true
	}
	elapsed := s.CreateTime - time.Now().Unix()
	if elapsed < expire {
		return true
	}
	return false
}
