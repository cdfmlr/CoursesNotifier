package accessToken

import (
	"fmt"
	"log"
	"time"
)

// accessTokenHolder 用来保持微信公众号的 access_token，所有需要使用 access_token 的地方应从此处获取.
//
// access_token 是公众号的全局唯一接口调用凭据，公众号调用各接口时都需使用access_token。
// 开发者需要进行妥善保存。
// access_token的存储至少要保留512个字符空间。
// access_token的有效期目前为2个小时，需定时刷新，重复获取将导致上次获取的access_token失效。
//
// 更多 access_token 说明详见：https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_access_token.html

type Holder struct {
	appID string
	appSecret string

	accessToken string
	createTime  int64
	expiresIn   int64
}

func NewHolder(appID string, appSecret string) *Holder {
	return &Holder{appID: appID, appSecret: appSecret}
}

func (h *Holder) isFresh() bool {
	if h.accessToken == "" {
		return false
	}
	elapsed := time.Now().Unix() - h.createTime
	return elapsed < h.expiresIn
}

func (h *Holder) refresh(current int, maxRetry int) {
	if current > maxRetry {
		log.Panic(fmt.Sprintf("Cannot fetch access_token after %d tries.", maxRetry))
		// TODO: Recover from this panic
	}
	log.Println("Try to get access token...")
	data, err := getAccessToken(h.appID, h.appSecret)
	if err != nil {
		log.Println(err)
	}
	if data.AccessToken == "" {
		log.Println("Failed to Get AccessToken, try again.")
		h.refresh(current+1, maxRetry)
	}
	h.accessToken = data.AccessToken
	h.createTime = time.Now().Unix()
	h.expiresIn = data.ExpiresIn
}

func (h *Holder) Get() string {
	if !h.isFresh() {
		h.refresh(0, 5)
	}
	return h.accessToken
}
