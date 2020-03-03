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

import (
	"log"
)

type GetCurrentTimeRespBody struct {
	Zc    int    `json:"zc"`
	STime string `json:"s_time"`
	ETime string `json:"e_time"`
	Xnxqh string `json:"xnxqh"`
}

/*
FetchCurrentTime 获取所提交的日期的时间、周次、学年等信息。

GET http://jwxt.xxxx.edu.cn/app.do?method=getCurrentTime&currDate={$查询日期}

Req:
	request.header{token:'运行身份验证authUser时获取到的token，有过期机制'},
	request.data{
		'method':'FetchCurrentTime',  //必填
		'currDate':  //格式为"YYYY-MM-DD"，必填，留空调用成功，但返回值均为null
	}

Resp:
	{
		"zc":20, //当前周次
		"e_time":"2019-01-20", //本周结束时间
		"s_time":"2019-01-14", //本周开始时间
		"xnxqh":"2018-2019-1" //学年学期名称
	}
*/
func GetCurrentTime(school, token, currDate string) (*GetCurrentTimeRespBody, error) {
	resp := &GetCurrentTimeRespBody{}
	q := map[string]string{
		"method":   "getCurrentTime",
		"currDate": currDate,
	}
	err := qzApiGet(school, token, resp, q)
	if err != nil {
		log.Println(err)
		return &GetCurrentTimeRespBody{}, err
	}
	return resp, nil
}
