package wxAccessToken

import (
	"log"
	"testing"
	"time"
)

func TestAccessTokenHolder_Get(t *testing.T) {
	h := NewHolder("***", "***")
	log.Println(h.Get())
	log.Println(h.expiresIn)
	timer := time.NewTimer(5 * time.Second)
	<-timer.C
	log.Println(time.Now().Unix() - h.createTime)
}
