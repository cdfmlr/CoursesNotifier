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

package wxPlatformServer

import (
	"example.com/CoursesNotifier/data"
	"fmt"
)

// Unsubscribe

type CoursesUnsubscribeSession struct {
	CoursesSerSession

	reqUser    string
	reqContent string
}

func NewCoursesUnsubscribeSession(reqUser string, reqContent string, databaseSource string) *CoursesUnsubscribeSession {
	s := &CoursesUnsubscribeSession{reqUser: reqUser, reqContent: reqContent}
	s.CoursesSerSession.databaseSource = databaseSource
	return s
}

// Verify 尝试退订课程提醒
func (s *CoursesUnsubscribeSession) Verify() string {

	s.GenerateVerification()

	return fmt.Sprintf(
		"(0x064) 您确认要退订课程提醒服务嘛T_T 若您意已决，请回复数字验证码：【%s】(五分钟内有效)",
		s.verification,
	)
}

// Continue 为用户办理课程提醒退订，
//  Continue 需要 Verify 提供的验证码
func (s *CoursesUnsubscribeSession) Continue(verificationCode string) string {
	if verificationCode != s.verification { // 验证码错误
		return "验证码错误，以为您取消退订。"
	}
	var totalAffected int64

	// 查询用户
	sdb := data.NewStudentDatabase(s.databaseSource)
	student, err := sdb.GetByWxUser(s.reqUser)

	if student.Sid == "" {
		return "没有查询到您到订阅，退订失败。"
	}

	affected, err := sdb.Delete(student.Sid)
	totalAffected += affected

	rdb := data.NewStudentCourseRelationshipDatabase(s.databaseSource)
	affected, err = rdb.DeleteBySid(student.Sid)
	totalAffected += affected

	if err != nil || totalAffected <= 0 {
		return "退订失败！"
	}

	return "退订成功！"
}
