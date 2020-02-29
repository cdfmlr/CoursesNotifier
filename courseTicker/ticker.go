package courseTicker

import (
	"log"
	"time"
)

// Ticker 在调用 Start 方法后将以 period 为周期，重复执行 task，直到 End 被调用
// 这个 Ticker 直接使用没有任何意义，只是周期性 log 打印 一串文字
// 建议 "继承" 它，并按照 TickerTask 接口 "重写" RunTickTask，在 RunTickTask 中实现具体的功能。
type Ticker struct {
	tickerId string
	period   time.Duration
	end      chan bool
	task     func()
}

// Start 设置 Ticker 从 time.Time 开始工作
func (t *Ticker) Start(time2Start time.Time) {
	time2Start = time2Start.Add(t.period * -1)
	if time2Start.Sub(time.Now()) > 0 {
		log.Printf("(Ticker {%s}) Sleep until time2Start: %v\n", t.tickerId, time2Start)
		timer := time.NewTimer(time2Start.Sub(time.Now()))
		<-timer.C
	}

	log.Printf("(Ticker {%s}) Begin to run TickTask periodically (period=%s).\n", t.tickerId, t.period)
	// 不断重复执行任务:
	go func() {
		for {
			select {
			case <-t.end:
				log.Printf("(Ticker {%s}) TickTask End Exed...\n", t.tickerId)
				return
			default:
				// 计算下一个执行时间
				now := time.Now()
				next := now.Add(t.period)
				// 等待到时间
				timer := time.NewTimer(next.Sub(now))
				<-timer.C
				// 执行任务
				t.RunTickTask()
			}
		}
	}()
}

// TickTask 为 Ticker 应做的周期性工作
func (t *Ticker) RunTickTask() {
	log.Printf("(Ticker {%s}) TickTask Run...\n", t.tickerId)
	if t.task != nil {
		t.task()
	}
}

// End 通知 TickTask 终止运行
func (t *Ticker) End() {
	log.Printf("(Ticker {%s}) Ending Ticker...\n", t.tickerId)
	t.end <- true
}
