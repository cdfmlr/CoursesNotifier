package wxPlatformServer

import (
	"fmt"
	"testing"
)

func TestCoursesSubscribeSession_Subscribe(t *testing.T) {
	s := NewCoursesSubscribeSession("reqUser", "订阅课表 201810000000 hd000000", "c:000123@/test?charset=utf8")
	//fmt.Println(s.SubscribeVerify())
	ss := s.Continue("2333")
	fmt.Println(ss)

	//s := NewCoursesSubscribeSession("reqUser", "订阅课表 201810000431 hd270516", "c:000123@/test?charset=utf8")
	//sr := s.SubscribeVerify()
	//fmt.Println(sr)
	//v := strings.Split(sr, "【")[1]
	//v = strings.Split(v, "】")[0]
	//fmt.Println(v)
	//ss := s.Subscribe(v)
	//fmt.Println(ss)
}
