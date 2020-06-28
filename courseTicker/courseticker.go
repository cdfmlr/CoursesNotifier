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

package courseTicker

import (
	"example.com/CoursesNotifier/util/ticker"
	"time"
)

// CoursesTicker 课程时钟
// ("继承" 自 Ticker，使用 Ticker 的 Start、RunTickTask、Stop 方法)
// 这个东西在调用 Start 方法后将以 period 为周期，重复执行 NotifyApproachingCourses 方法，直到 Stop 被调用。
type CoursesTicker struct {
	ticker.Ticker
	databaseSource             string
	minuteBeforeCourseToNotify float64
	notifiers                  []Notifier
}

func NewCoursesTicker(tickerId string, databaseSource string, period time.Duration, minuteBeforeCourseToNotify float64, notifiers []Notifier) *CoursesTicker {
	ct := &CoursesTicker{databaseSource: databaseSource, minuteBeforeCourseToNotify: minuteBeforeCourseToNotify, notifiers: notifiers}
	// 初始化 "父类"
	ct.TickerId = tickerId
	ct.Period = period
	ct.End = make(chan bool)
	ct.Task = ct.NotifyApproachingCourses
	return ct
}
