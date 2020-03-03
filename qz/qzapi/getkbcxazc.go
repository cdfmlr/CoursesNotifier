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

package qzapi

import "log"

type GetKbcxAzcRespBodyItem struct {
	Jsxm string `json:"jsxm"`
	Jsmc string `json:"jsmc"`
	Kssj string `json:"kssj"`
	Jssj string `json:"jssj"`
	Kkzc string `json:"kkzc"`
	Kcsj string `json:"kcsj"`
	Kcmc string `json:"kcmc"`
	Sjbz string `json:"sjbz"`
}

/*
GetKbcxAzc 获取一周的课程信息

GET http://jwxt.xxxx.edu.cn/app.do?method=getKbcxAzc&xh={$学号}&xnxqid={$学年学期ID}&zc={$周次}

Req:
	request.header{token:'运行身份验证authUser时获取到的token，有过期机制'},
	request.data{
		'method':'GetKbcxAzc',  //必填
		'xh':'2017168xxxxx',  //必填，使用与获取token时不同的学号，则可以获取到新输入的学号的课表
		'xnxqid':'2018-2019-1',  //格式为"YYYY-YYYY-X"，非必填，不包含时返回当前日期所在学期课表
		'zc':'1'  //必填
	}

Resp:
	[
		{
			"jsxm":"张三", //教师姓名
			"jsmc":"教学楼101", //教室名称
			"jssj":"10:00", //结束时间
			"kssj":"08:00", //开始时间
			"kkzc":"1", //开课周次，有三种已知格式1)a-b、2)a,b,c、3)a-b,c-d
			"kcsj":"10506", //课程时间，格式x0a0b，意为星期x的第a,b节上课
			"kcmc":"大学英语", //课程名称
			"sjbz":"0" //具体意义未知，据观察值为1时本课单周上，2时双周上
		},{
			"jsxm":"李四",
			"jsmc":"教学楼101",
			"jssj":"12:00",
			"kssj":"10:00",
			"kkzc":"1",
			"kcsj":"1000000",
			"kcmc":"微积分",
			"sjbz":"0"
		}
	]
*/
func GetKbcxAzc(school, token, xh, xnxqid, zc string) ([]GetKbcxAzcRespBodyItem, error) {
	var resp []GetKbcxAzcRespBodyItem
	q := map[string]string{
		"method": "getKbcxAzc",
		"xh":     xh,
		"xnxqid": xnxqid,
		"zc":     zc,
	}
	err := qzApiGet(school, token, &resp, q)
	if err != nil {
		log.Println(err)
		return []GetKbcxAzcRespBodyItem{}, err
	}
	return resp, nil
}
