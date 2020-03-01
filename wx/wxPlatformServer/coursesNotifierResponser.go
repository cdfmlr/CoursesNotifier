package wxPlatformServer

import (
	"regexp"
	"strings"
)

type CourseNotifierResponser struct {
	sessionMap     map[string]VerifySerSession // {reqUser(wxUser): coursesSerSession}
	databaseSource string
}

// 订阅课表首先是粗略判断用户输入是否合法，
// 然后尝试拿用户的输入登录强智系统，
// 如果登录成功，则返回真实姓名、系、一个验证码给用户，问他正不正确、要不要办
// 同时把这个强智客户端搁一下，等待用户返回验证码，
// 如果这时接收到一条消息是之前的用户发的，同时内容是刚才那个验证码，就给他查课表、写入库，告诉他服务开了
//
// 退订也差不多这个流程：判断 -> 尝试 -> 验证码 -> 写库

func (c CourseNotifierResponser) Do(reqUser string, reqContent string) (respContent string) {
	reqContent = strings.TrimSpace(reqContent) // 去掉首位空白字符
	switch {
	case isReqSubscribe(reqContent):
		c.sessionMap[reqUser] = NewCoursesSubscribeSession(reqUser, reqContent, c.databaseSource)
		return c.sessionMap[reqUser].Verify()
	case isReqUnsubscribe(reqContent):
		c.sessionMap[reqUser] = NewCoursesUnsubscribeSession(reqUser, reqContent, c.databaseSource)
		return c.sessionMap[reqUser].Verify()
	case isReqVerification(reqContent):
		if c.sessionMap[reqUser] != nil {
			return c.sessionMap[reqUser].Continue(reqContent)
		} else {
			return "😯你发这个干嘛？"
		}
	}
	return `欢迎使用 NCEPU(Baoding) 课程提醒系统！

订阅课程提醒业务，请回复"订阅课表" + 空格 + 学号 + 空格 + 教务密码
例如："订阅课表 209910000999 hd666666"（不输入引号）；

退订课程提醒业务，请回复"退订"二字。

(本服务非官方提供，对服务质量不做保证！)
All rights reserved © 2020 CDFMLR
`
}

// isReqSubscribe 判断请求是否为**订阅**操作，是则返回 true，否则 false
// 订阅操作请求内容格式如下：
// 		"订阅课表 201810000999 hd666666"
// 即需符合 "订阅课表" + 空格 + 学号 + 空格 + 教务密码
func isReqSubscribe(reqContent string) bool {
	rs := strings.Split(reqContent, " ")
	if len(rs) == 3 && rs[0] == "订阅课表" { // 符合订阅操作格式
		matched, _ := regexp.MatchString(`^\d{12}$$`, rs[1]) // 学号是数字, 且长度正常
		return matched
	}
	return false
}

// isReqSubscribe 判断请求是否为**退订**操作，是则返回 true，否则 false
// 退订操作应该是：
//		退订
//	这两个字。
func isReqUnsubscribe(reqContent string) bool {
	return reqContent == "退订"
}

// isReqVerification 判断请求是否为**验证码**，是则返回 true，否则 false
// 验证码应该是四位随机数字，形如：
//		6982
func isReqVerification(reqContent string) bool {
	matched, _ := regexp.MatchString(`^\d{4}$$`, reqContent)
	return matched
}
