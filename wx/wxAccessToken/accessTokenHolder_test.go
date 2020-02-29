package wxAccessToken

import (
	"log"
	"testing"
	"time"
)

func TestAccessTokenHolder_Get(t *testing.T) {
	h := NewHolder("wx63cf76ed67d69bb1", "8a62c82aeac97ebf79b4617049499302")
	log.Println(h.Get())
	log.Println(h.expiresIn)
	timer := time.NewTimer(5 * time.Second)
	<-timer.C
	log.Println(time.Now().Unix() - h.createTime)
}