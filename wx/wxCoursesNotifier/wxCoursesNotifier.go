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

package wxCoursesNotifier

import (
	"bytes"
	"encoding/json"
	"example.com/CoursesNotifier/models"
	"example.com/CoursesNotifier/util/briefBullshitGenerator"
	"example.com/CoursesNotifier/wx/wxAccessToken"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// WxNotifier 通过微信发送上课提醒
type WxNotifier struct {
	courseNoticeTemplateID string
	wxTokenHolder          *wxAccessToken.Holder
	bullshitDataFilePath   string
}

func New(courseNoticeTemplateID string, wxTokenHolder *wxAccessToken.Holder, bullshitDataFilePath string) *WxNotifier {
	return &WxNotifier{courseNoticeTemplateID: courseNoticeTemplateID, wxTokenHolder: wxTokenHolder, bullshitDataFilePath: bullshitDataFilePath}
}

func (w WxNotifier) Notify(student *models.Student, course *models.Course) {
	bullshit := briefBullshitGenerator.Generate(w.bullshitDataFilePath)
	noticeBody, err := w.makeCourseNoticeBody(student.WxUser, course.Name, course.Location, course.Teacher, course.Begin, course.End, course.Week, bullshit)
	err = w.postCourseNotify(noticeBody)
	if err != nil {
		// TODO: Do something here.
	}
}

// NoticeItem, CourseData, WxNotice 为微信公众号课程通知 json 的各级组成部分
// 对应的微信模版内容应设置如下：
// {{first.DATA}} 课程：{{course.DATA}} 地点：{{location.DATA}} 老师：{{teacher.DATA}} 时间：{{time.DATA}} 教学周：{{week.DATA}} --- {{bullshit.DATA}} {{remark.DATA}}
type NoticeItem struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

type CourseData struct {
	First    NoticeItem `json:"first"`
	Course   NoticeItem `json:"course"`
	Location NoticeItem `json:"location"`
	Teacher  NoticeItem `json:"teacher"`
	BETime   NoticeItem `json:"time"`
	Week     NoticeItem `json:"week"`
	Bullshit NoticeItem `json:"bullshit"`
	Remark   NoticeItem `json:"remark"`
}

type WxNotice struct {
	ToUser     string     `json:"touser"`
	TemplateId string     `json:"template_id"`
	Data       CourseData `json:"data"`
}

// makeCourseNoticeBody 构建微信上课通知 json
func (w WxNotifier) makeCourseNoticeBody(toUser, course, location, teacher, begin, end, week, bullshit string) ([]byte, error) {
	notice := WxNotice{
		ToUser:     toUser,
		TemplateId: w.courseNoticeTemplateID,
		Data: CourseData{
			First: NoticeItem{
				Value: "滚去上课" + "\n",
				Color: "#e51c23",
			},
			Course: NoticeItem{
				Value: course + "\n",
				Color: "#173177",
			},
			Location: NoticeItem{
				Value: location + "\n",
				Color: "#173177",
			},
			Teacher: NoticeItem{
				Value: teacher + "\n",
				Color: "#173177",
			},
			BETime: NoticeItem{
				Value: begin + "~" + end + "\n",
				Color: "#173177",
			},
			Week: NoticeItem{
				Value: week + "\n",
				Color: "#173177",
			},
			Bullshit: NoticeItem{
				Value: bullshit,
				Color: "#5677fc",
			},
			Remark: NoticeItem{
				Value: "\n但还是要好好听课哦💪" + "\n\n",
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
