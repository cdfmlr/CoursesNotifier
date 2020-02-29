package courseTicker

import (
	"example.com/CoursesNotifier/wx/wxAccessToken"
	"example.com/CoursesNotifier/wx/wxCoursesNotifier"
	"testing"
	"time"
)

func TestCoursesTicker(t *testing.T) {
	h := wxAccessToken.NewHolder("wx63cf76ed67d69bb1", "8a62c82aeac97ebf79b4617049499302")
	ct := NewCoursesTicker(
		"courseTicker", "c:000123@/test?charset=utf8", time.Second*2, 10,
		[]Notifier{LogNotifier("LN"), wxCoursesNotifier.NewWxNotifier("-mQRq0B8nP5g5auWBOaAO0uLr64owmoNtwqJBrkz5G0", h)},
	)
	ct.Start(time.Now())
	timer := time.NewTimer(time.Second * 10)
	<-timer.C
	ct.End()
	timer = time.NewTimer(time.Second * 10)
	<-timer.C
}
