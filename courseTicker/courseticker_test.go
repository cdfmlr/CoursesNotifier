package courseTicker

import (
	"testing"
	"time"
)

func TestCoursesTicker(t *testing.T) {
	ct := NewCoursesTicker("courseTicker", "c:000123@/test?charset=utf8", time.Second*2, 10, []Notifier{LogNotifier("LN")})
	ct.Start(time.Now())
	timer := time.NewTimer(time.Second * 10)
	<-timer.C
	ct.End()
	timer = time.NewTimer(time.Second * 10)
	<-timer.C
}
