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

package models

import (
	"time"
)

type Student struct {
	Sid         string   // 学号
	Pwd         string   // 教务密码
	WxUser      string   // 微信用户公众号openid
	Courses     []Course // 课程列表
	CreateTime  int64    // 创建时间
	ExamResults string   // 已出成绩的课程名字列表
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
