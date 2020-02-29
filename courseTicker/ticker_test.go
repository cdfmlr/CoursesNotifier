package courseTicker

import (
	"testing"
	"time"
)

func TestTicker(t1 *testing.T) {
	t := Ticker{tickerId: "ticker", period: time.Second * 3, end: make(chan bool)}
	t.Start(time.Now())
	timer := time.NewTimer(time.Second * 10)
	<-timer.C
	t.End()
	timer = time.NewTimer(time.Second * 5)
	<-timer.C
}