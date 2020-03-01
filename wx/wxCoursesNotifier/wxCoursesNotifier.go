package wxCoursesNotifier

import (
	"bytes"
	"encoding/json"
	"example.com/CoursesNotifier/models"
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
}

func New(courseNoticeTemplateID string, wxTokenHolder *wxAccessToken.Holder) *WxNotifier {
	return &WxNotifier{courseNoticeTemplateID: courseNoticeTemplateID, wxTokenHolder: wxTokenHolder}
}

func (w WxNotifier) Notify(student *models.Student, course *models.Course) {
	noticeBody, err := w.makeCourseNoticeBody(student.WxUser, course.Name, course.Location, course.Teacher, course.Week)
	err = w.postCourseNotify(noticeBody)
	if err != nil {
		// TODO: Do something here.
	}
}

// NoticeItem, CourseData, WxNotice 为微信公众号课程通知 json 的各级组成部分
type NoticeItem struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

type CourseData struct {
	First    NoticeItem `json:"first"`
	Course   NoticeItem `json:"course"`
	Location NoticeItem `json:"location"`
	Teacher  NoticeItem `json:"teacher"`
	Week     NoticeItem `json:"week"`
	Remark   NoticeItem `json:"remark"`
}

type WxNotice struct {
	ToUser     string     `json:"touser"`
	TemplateId string     `json:"template_id"`
	Data       CourseData `json:"data"`
}

// makeCourseNoticeBody 构建微信上课通知 json
func (w WxNotifier) makeCourseNoticeBody(toUser, course, location, teacher, week string) ([]byte, error) {
	notice := WxNotice{
		ToUser:     toUser,
		TemplateId: w.courseNoticeTemplateID,
		Data: CourseData{
			First: NoticeItem{
				Value: "滚去上课" + "\n\n",
				Color: "#173177",
			},
			Course: NoticeItem{
				Value: course + "\n\n",
				Color: "#173177",
			},
			Location: NoticeItem{
				Value: location + "\n\n",
				Color: "#173177",
			},
			Teacher: NoticeItem{
				Value: teacher + "\n\n",
				Color: "#173177",
			},
			Week: NoticeItem{
				Value: week + "\n\n",
				Color: "#173177",
			},
			Remark: NoticeItem{
				Value: "🤯" + "\n\n",
				Color: "#173177",
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
