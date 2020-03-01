package wxPlatformServer

import (
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
)

type WxPlatformServer struct {
	token          string
	responser      Responser
	databaseSource string
}

func (s *WxPlatformServer) Handle(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if !s.validateWxRequest(w, r) {
		log.Println("WeChat Service: this http request is not from Wechat platform!")
		return
	}
	if r.Method == "POST" {
		textRequestBody := s.parseTextRequestBody(r)
		if textRequestBody != nil {
			thisSer := textRequestBody.ToUserName

			reqUser := textRequestBody.FromUserName
			reqContent := textRequestBody.Content

			respContent := s.responser.Do(reqUser, reqContent)

			responseTextBody, err := s.makeTextResponseBody(thisSer, reqUser, respContent)

			_, err = fmt.Fprint(w, string(responseTextBody))

			if err != nil {
				log.Println("WeChat Service: resp error: ", err)
			}
		}

		if err != nil {
			log.Println("WeChat Service: req error: ", err)
			return
		}
		// fmt.Fprint(w, string(responseTextBody))
	}
}

func (s *WxPlatformServer) validateWxRequest(w http.ResponseWriter, r *http.Request) bool {
	r.ParseForm()

	signature := r.FormValue("signature")

	timestamp := r.FormValue("timestamp")
	nonce := r.FormValue("nonce")

	echostr := r.FormValue("echostr")

	hashcode := s.makeSignature(s.token, timestamp, nonce)

	log.Printf("Try validateWxRequest: hashcode: %s, signature: %s\n", hashcode, signature)
	if hashcode == signature {
		fmt.Fprintf(w, "%s", echostr)
		return true
	} else {
		fmt.Fprintf(w, "hashcode != signature")
	}
	return false
}

func (s *WxPlatformServer) makeSignature(token, timestamp, nonce string) string {
	sl := []string{token, timestamp, nonce}
	sort.Strings(sl)

	sha := sha1.New()
	io.WriteString(sha, strings.Join(sl, ""))

	return fmt.Sprintf("%x", sha.Sum(nil))
}

type TextRequestBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Content      string
	MsgId        int
}

type TextResponseBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATAText
	FromUserName CDATAText
	CreateTime   time.Duration
	MsgType      CDATAText
	Content      CDATAText
}

type CDATAText struct {
	Text string `xml:",innerxml"`
}

func value2CDATA(v string) CDATAText {
	return CDATAText{"<![CDATA[" + v + "]]>"}
}

func (s *WxPlatformServer) parseTextRequestBody(r *http.Request) *TextRequestBody {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	fmt.Println(string(body))
	requestBody := &TextRequestBody{}
	xml.Unmarshal(body, requestBody)
	return requestBody
}

func (s *WxPlatformServer) makeTextResponseBody(fromUserName, toUserName, content string) ([]byte, error) {
	textResponseBody := &TextResponseBody{}
	textResponseBody.FromUserName = value2CDATA(fromUserName)
	textResponseBody.ToUserName = value2CDATA(toUserName)
	textResponseBody.MsgType = value2CDATA("text")
	textResponseBody.Content = value2CDATA(content)
	textResponseBody.CreateTime = time.Duration(time.Now().Unix())
	return xml.MarshalIndent(textResponseBody, " ", "  ")
}
