/*
 * Copyright 2020 CDFMLR
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package wxExamResultNotifier

import (
	"bytes"
	"encoding/json"
	"example.com/CoursesNotifier/models"
	"example.com/CoursesNotifier/util/briefBullshitGenerator"
	"example.com/CoursesNotifier/wx/wxAccessToken"
	"fmt"
	"github.com/cdfmlr/qzgo"
	"io/ioutil"
	"log"
	"net/http"
)

// WxNotifier 通过微信发送上课提醒
type WxNotifier struct {
	examResultNoticeTemplateID string
	wxTokenHolder              *wxAccessToken.Holder
	bullshitDataFilePath       string
}

func New(examResultNoticeTemplateID string, wxTokenHolder *wxAccessToken.Holder, bullshitDataFilePath string) *WxNotifier {
	return &WxNotifier{examResultNoticeTemplateID: examResultNoticeTemplateID, wxTokenHolder: wxTokenHolder, bullshitDataFilePath: bullshitDataFilePath}
}

func (w WxNotifier) Notify(student *models.Student, result qzgo.GetCjcxRespBodyItem) {
	bullshit := briefBullshitGenerator.Generate(w.bullshitDataFilePath)
	noticeBody, err := w.makeExamResultNoticeBody(student.WxUser, result, bullshit)
	err = w.postCourseNotify(noticeBody)
	if err != nil {
		// TODO: Do something here.
	}
}

// NoticeItem, ExamResultData, WxNotice 为微信公众号考试成绩通知 json 的各级组成部分
// 对应的微信模版内容应设置如下：
// 		{{first.DATA}}
//		课程：{{course.DATA}}
//		成绩：{{result.DATA}}
//		学分：{{xf.DATA}}
//		课程类别：{{ctype.DATA}}
//		考试性质：{{examtype.DATA}}
//		{{remark.DATA}}
type NoticeItem struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

type ExamResultData struct {
	First    NoticeItem `json:"first"`
	Course   NoticeItem `json:"course"`
	Result   NoticeItem `json:"result"`
	Xf       NoticeItem `json:"xf"`
	Ctype    NoticeItem `json:"ctype"`
	Examtype NoticeItem `json:"examtype"`
	Remark   NoticeItem `json:"remark"`
}

type WxNotice struct {
	ToUser     string         `json:"touser"`
	TemplateId string         `json:"template_id"`
	Data       ExamResultData `json:"data"`
}

// makeCourseNoticeBody 构建微信上课通知 json
func (w WxNotifier) makeExamResultNoticeBody(toUser string, result qzgo.GetCjcxRespBodyItem, bullshit string) ([]byte, error) {
	notice := WxNotice{
		ToUser:     toUser,
		TemplateId: w.examResultNoticeTemplateID,
		Data: ExamResultData{
			First: NoticeItem{
				Value: "出成绩啦😱" + "\n",
				Color: "#e51c23",
			},
			Course: NoticeItem{
				Value: result.Kcmc + "\n",
				Color: "#173177",
			},
			Result: NoticeItem{
				Value: result.Zcj + "\n",
				Color: "#173177",
			},
			Xf: NoticeItem{
				Value: fmt.Sprintf("%v", result.Xf) + "\n",
				Color: "#173177",
			},
			Ctype: NoticeItem{
				Value: result.Kclbmc + "\n",
				Color: "#173177",
			},
			Examtype: NoticeItem{
				Value: result.Ksxzmc + "\n",
				Color: "#173177",
			},
			Remark: NoticeItem{
				Value: bullshit + "\n\n",
				Color: "#000000",
			},
		},
	}
	return json.MarshalIndent(notice, " ", "  ")
}

// postCourseNotify 发送微信公众号上课通知请求
func (w WxNotifier) postCourseNotify(CourseNoticeBody []byte) error {
	// http请求方式: POST https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=ACCESS_TOKEN
	url := fmt.Sprintf(
		"https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s",
		w.wxTokenHolder.Get(),
	)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(CourseNoticeBody))
	if err != nil {
		log.Println("postCourseNotify Failed:", err)
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		err = NotifyFailed(fmt.Sprintf("postCourseNotify Failed: \n\t"+
			"|--> response Status: %s \n\t"+
			"|--> response Header: %s\n\t"+
			"|--> response Body: %s\n",
			resp.Status,
			resp.Header,
			string(body),
		))
		log.Println(err)
		return err
	}

	return nil
}

// NotifyFailed 请求返回状态值不为200时抛出的错误
type NotifyFailed string

func (n NotifyFailed) Error() string {
	return string(n)
}
