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