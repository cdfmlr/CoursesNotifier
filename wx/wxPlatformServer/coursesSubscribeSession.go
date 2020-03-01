package wxPlatformServer

import (
	"example.com/CoursesNotifier/models"
	"example.com/CoursesNotifier/qzclient"
	"fmt"
	"strings"
)

type CoursesSubscribeSession struct {
	CoursesSerSession

	reqUser    string
	reqContent string

	qzClient     *qzclient.Client
}

func NewCoursesSubscribeSession(reqUser string, reqContent string, databaseSource string) *CoursesSubscribeSession {
	s := &CoursesSubscribeSession{reqUser: reqUser, reqContent: reqContent}
	s.CoursesSerSession.databaseSource = databaseSource
	return s
}

// Verify 尝试拿用户请求中的信息登录强智系统，检测是否具有办理订阅课表的资格
// 若登录强智系统成功，即用户拥有订阅资格，这是返回强智系统中用户真实姓名、院系、以及一个验证码给用户
//
// 订阅操作请求内容格式如下：
// 		"订阅课表 201810000999 hd666666"
// 即需符合 "订阅课表" + 空格 + 学号 + 空格 + 教务密码
func (s *CoursesSubscribeSession) Verify() string {
	rs := strings.Split(s.reqContent, " ")
	sid, pwd := rs[1], rs[2]
	student := models.NewStudent(sid, pwd, s.reqUser)

	s.qzClient = qzclient.New(*student)
	authRespBody, err := s.qzClient.AuthUser()
	realName, school := authRespBody.UserRealName, authRespBody.UserDwmc // 姓名、院系

	if err != nil {
		return "(0x196) 抱歉，系统无法处理您提供的信息，请查正后再试。若问题持续存在，请联系管理员。"
	}

	ch := make(chan []models.Course)
	err = s.qzClient.FetchCurrentTime()
	go s.qzClient.FetchAllTermCourses(ch)

	// 去重提取出课程名称、老师
	courses := make(map[string]string)
	for _, c := range <-ch {
		courses[c.Name] = c.Teacher
	}

	if err != nil || len(courses) == 0 {
		return "(0x196) 抱歉，系统无法获取您的课表，请查正后再试。若问题持续存在，请联系管理员。"
	}

	// 合并为一个可读字符串
	coursesStr := ""
	for c, t := range courses {
		coursesStr = fmt.Sprintf("%s\n%s (%s);", coursesStr, c, t)
	}

	s.GenerateVerification()

	return fmt.Sprintf(
		"(0x064) 根据您提供的信息，我们查询到您是 %s 的 %s。您本学期的课程有: %s\n如果信息正确，且确认订阅课程提醒服务，请回复数字验证码：【%s】",
		school,
		realName,
		coursesStr,
		s.verification,
	)
}

// Continue 为用户办理课程提醒登记，
//  Continue 需要 Verify 提供的验证码
func (s *CoursesSubscribeSession) Continue(verificationCode string) string {
	if verificationCode != s.verification { // 验证码错误
		return "验证码错误，以为您取消订阅。"
	}
	affected := s.qzClient.Save(s.databaseSource)
	if affected > 0 {
		return "(0x0C8) 订阅成功！"
	} else { // 数据库一行都没动，其实是失败的！
		return "(0x130) 订阅成功！"
	}
}
