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

package qzapi

import "time"

// 和课程表中
//		"kcsj":"10506", //课程时间，格式x0a0b，意为星期x的第a,b节上课
// 对应的 x
//
// NOTE: 不确定周六、周日的表示形式，我猜测是 6、7

type QzWeekday int

const (
	_ QzWeekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

// TimeWeekToQzWeekday 将 time.Weekday 转化为 qzWeekday
func TimeWeekToQzWeekday(weekday time.Weekday) QzWeekday {
	if weekday == time.Sunday {
		return Sunday
	}
	return QzWeekday(weekday)
}