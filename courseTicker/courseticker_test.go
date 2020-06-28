package courseTicker

import (
	"example.com/CoursesNotifier/wx/wxAccessToken"
	"example.com/CoursesNotifier/wx/wxCoursesNotifier"
	"testing"
	"time"
)

func TestCoursesTicker(t *testing.T) {
	h := wxAccessToken.NewHolder("***", "***")
	ct := NewCoursesTicker(
		"courseTicker", "c:***@/test?charset=utf8", time.Second*2, 10,
		[]Notifier{LogNotifier("LN"), wxCoursesNotifier.New("-***", h, "")},
	)
	ct.Start(time.Now())
	timer := time.NewTimer(time.Second * 10)
	<-timer.C
	ct.Stop()
	timer = time.NewTimer(time.Second * 10)
	<-timer.C
}
