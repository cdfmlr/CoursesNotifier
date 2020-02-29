package courseTicker

import "time"

type CoursesTicker struct {
	databaseSource             string
	duration                   time.Duration
	minuteBeforeCourseToNotify int
}
