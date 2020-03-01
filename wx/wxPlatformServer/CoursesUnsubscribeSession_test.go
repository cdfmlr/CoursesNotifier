package wxPlatformServer

import (
	"fmt"
	"strings"
	"testing"
)

func TestCoursesUnsubscribeSession(t *testing.T) {
	s := NewCoursesUnsubscribeSession("hahaha", "退订", "c:000123@/test?charset=utf8")

	ss := s.Continue("2333")
	fmt.Println(ss)

	sr := s.Verify()
	fmt.Println(sr)
	v := strings.Split(sr, "【")[1]
	v = strings.Split(v, "】")[0]
	fmt.Println(v)
	sss := s.Continue(v)
	fmt.Println(sss)
}
