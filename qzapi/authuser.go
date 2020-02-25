package qzapi

import "log"

type AuthUserRespBody struct {
	Flag         string `json:"flag"`
	UserRealName string `json:"userrealname"`
	Token        string `json:"token"`
	UserDwmc     string `json:"userdwmc"`
	UserType     string `json:"usertype"`
	Msg          string `json:"msg"`
}

/*
authUser 登录帐号。

在调用其它接口前，需要使用本接口获取token。并在调用其它接口的请求的header中添加本接口获得的token参数。

GET http://jwxt.xxxx.edu.cn/app.do?method=authUser&xh={$学号}&pwd={$密码}

Req:
	{
		"method":'authUser',  //必填
		"xh":'登陆教务系统使用的学号',  //必填
		"pwd":'登陆教务系统需要的密码'  //必填
	}

Resp:
	{
		"flag":"1", //是否成功 #成功1 失败0
		"userrealname":"张三", //用户真实姓名 #失败 null
		"token":"", //令牌 #失败 -1
		"userdwmc":"XXXX学院", //用户所在学院名称 #失败 null
		"usertype":"2", //用户类别 #已知学生身份为2 失败 null
		"msg":"登录成功" //返回消息
	}
*/
func authUser(school, xh, pwd string) (*AuthUserRespBody, error) {
	authUserRespBody := &AuthUserRespBody{}
	q := map[string]string{
		"method": "authUser",
		"xh":     xh,
		"pwd":    pwd,
	}
	err := qzApiGet(school, "", authUserRespBody, q)
	if err != nil {
		log.Println(err)
		return &AuthUserRespBody{}, err
	}
	return authUserRespBody, nil
}
