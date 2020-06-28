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

package examResultTicker

import (
	"example.com/CoursesNotifier/data"
	"example.com/CoursesNotifier/util/ticker"
	"github.com/cdfmlr/qzgo"
	"log"
	"math/rand"
	"strings"
	"time"
)

const SCHOOL = "ncepu"

// ExamResultTicker 课程时钟
// ("继承" 自 Ticker，使用 Ticker 的 Start、RunTickTask、Stop 方法)
// 这个东西在调用 Start 方法后将以 period 为周期，重复执行 NotifyApproachingCourses 方法，直到 Stop 被调用。
type ExamResultTicker struct {
	ticker.Ticker
	databaseSource string
	notifiers      []Notifier
}

func NewExamResultTicker(tickerId string, databaseSource string, period time.Duration, notifiers []Notifier) *ExamResultTicker {
	ert := &ExamResultTicker{databaseSource: databaseSource, notifiers: notifiers}
	// 初始化 "父类"
	ert.TickerId = tickerId
	ert.Period = period
	ert.End = make(chan bool)
	ert.Task = ert.notifyNewExamResult
	return ert
}

// 通知新获取到的成绩
func (e *ExamResultTicker) notifyNewExamResult() {
	// Get all students
	sdb := data.NewStudentDatabase(e.databaseSource)
	students, err := sdb.GetStudents()
	if err != nil {
		log.Println(err)
		return
	}
	if len(students) <= 0 {
		log.Println("no student")
		return
	}

	// pick a random student to new a qzgo.Client
	randStudent := students[rand.Intn(len(students))]

	client, err := qzgo.NewClientLogin(SCHOOL, randStudent.Sid, randStudent.Pwd)
	if err != nil {
		log.Println(err)
		return
	}

	// Query Exam Results for all students
	for _, s := range students {
		// 随机睡一会儿在搞，别反爬虫给踢了
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		cjcxResp, err := client.GetCjcx(s.Sid, "")
		if err != nil {
			log.Println(err)
			continue
		}
		for _, cj := range cjcxResp.Result {
			if !strings.Contains(s.ExamResults, cj.Kcmc) { // 新出的成绩
				// 通知
				for _, n := range e.notifiers {
					n.Notify(&s, cj)
				}
				// 保存
				s.ExamResults += cj.Kcmc + ","
				affected, err := sdb.Update(s.Sid, s)
				if affected != 1 || err != nil {
					log.Println("notifyNewExamResult update sdb unexpected: rowsAffected =", affected, ", err =", err)
				}
			}
		}
	}

}
