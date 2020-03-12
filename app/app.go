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

package app

import (
	"example.com/CoursesNotifier/courseTicker"
	"example.com/CoursesNotifier/util/jsonFileLoader"
	"example.com/CoursesNotifier/wx/wxAccessToken"
	"example.com/CoursesNotifier/wx/wxCoursesNotifier"
	"example.com/CoursesNotifier/wx/wxPlatformServer"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

type App struct {
	conf    AppConf
	runtime AppRuntime
}

type AppConf struct {
	Wx     WxConf     `json:"wx"`
	Ticker TickerConf `json:"ticker"`
	Data   DataConf   `json:"data"`
}

type WxConf struct {
	AppID                  string `json:"app_id"`
	AppSecret              string `json:"app_secret"`
	ReqToken               string `json:"req_token"`
	CourseNoticeTemplateID string `json:"course_notice_template_id"`
}

type TickerConf struct {
	TimeToStart                string  `json:"time_to_start"`
	PeriodMinute               int     `json:"period_minute"`
	MinuteBeforeCourseToNotify float64 `json:"minute_before_course_to_notify"`
}

type DataConf struct {
	Database         string `json:"database"`
	BullshitDataFile string `json:"bullshit_data_file"`
}

type AppRuntime struct {
	pWxAccessTokenHolder *wxAccessToken.Holder
	pCoursesTicker       *courseTicker.CoursesTicker
	pWxPlatformServer    *wxPlatformServer.WxPlatformServer
}

func New(configFilePath string) *App {
	a := &App{}
	// 读取配置文件
	err := jsonFileLoader.Load(configFilePath, &a.conf)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to load configuration: ", err)
	}

	return a
}

// Test 测试配置完整性、正确行, 若配置完整、可用，则返回 nil，否则返回错误 error
func (app *App) Test() error {
	errs := make([]*ConfingMissing, 0)
	// fmt.Println(app)
	if app.conf.Ticker.PeriodMinute == int(0) {
		errs = append(errs, NewConfingMissing("Ticker.PeriodMinute"))
	}
	if strings.TrimSpace(app.conf.Ticker.TimeToStart) == "" {
		errs = append(errs, NewConfingMissing("Ticker.TimeToStart"))
	}
	if app.conf.Ticker.MinuteBeforeCourseToNotify == float64(0) {
		errs = append(errs, NewConfingMissing("Ticker.MinuteBeforeCourseToNotify"))
	}
	if strings.TrimSpace(app.conf.Wx.AppID) == "" {
		errs = append(errs, NewConfingMissing("Wx.AppID"))
	}
	if strings.TrimSpace(app.conf.Wx.AppSecret) == "" {
		errs = append(errs, NewConfingMissing("Wx.AppSecret"))
	}
	if strings.TrimSpace(app.conf.Wx.CourseNoticeTemplateID) == "" {
		errs = append(errs, NewConfingMissing("Wx.CourseNoticeTemplateID"))
	}
	if strings.TrimSpace(app.conf.Wx.ReqToken) == "" {
		errs = append(errs, NewConfingMissing("Wx.ReqToken"))
	}
	if strings.TrimSpace(app.conf.Data.Database) == "" {
		errs = append(errs, NewConfingMissing("Data.Database"))
	}
	if strings.TrimSpace(app.conf.Data.BullshitDataFile) == "" {
		errs = append(errs, NewConfingMissing("Data.BullshitDataFile"))
	}
	if len(errs) != 0 {
		s := ""
		for _, e := range errs {
			s += e.miss + ", "
		}
		return *NewConfingMissing(strings.Trim(s, ", "))
	}
	return nil
}

func (app *App) Run() {
	// 初始化运行时组件
	app.initWxAccessTokenHolder()
	app.initWxPlatformServer()
	app.initCoursesTicker()

	// 启动守护任务
	app.runWxPlatformServer()
	app.runCourseTicker()
}

// 初始化微信公众号 access_token holder
func (app *App) initWxAccessTokenHolder() {
	app.runtime.pWxAccessTokenHolder = wxAccessToken.NewHolder(app.conf.Wx.AppID, app.conf.Wx.AppSecret)
}

// 初始化微信公众号服务
func (app *App) initWxPlatformServer() {
	responser := wxPlatformServer.NewCourseNotifierResponser(app.conf.Data.Database)
	app.runtime.pWxPlatformServer = wxPlatformServer.New(app.conf.Wx.ReqToken, responser, app.conf.Data.Database)
}

// 初始化课程时钟
func (app *App) initCoursesTicker() {
	// 新建 课程时钟
	app.runtime.pCoursesTicker = courseTicker.NewCoursesTicker(
		"CourseTicker",
		app.conf.Data.Database,
		time.Duration(app.conf.Ticker.PeriodMinute)*time.Minute,
		app.conf.Ticker.MinuteBeforeCourseToNotify,
		[]courseTicker.Notifier{
			courseTicker.LogNotifier("LogNotifier"),
			wxCoursesNotifier.New(
				app.conf.Wx.CourseNoticeTemplateID,
				app.runtime.pWxAccessTokenHolder,
				app.conf.Data.BullshitDataFile,
			),
		},
	)
}

// 开始运行微信服务
func (app *App) runWxPlatformServer() {
	http.HandleFunc("/wx", app.runtime.pWxPlatformServer.Handle)
}

// 开始运行课程时钟
func (app *App) runCourseTicker() {
	timeToStart, _ := time.Parse("2006-01-02 15:04", app.conf.Ticker.TimeToStart)
	app.runtime.pCoursesTicker.Start(timeToStart)
}

type ConfingMissing struct {
	miss string
}

func NewConfingMissing(miss string) *ConfingMissing {
	return &ConfingMissing{miss: miss}
}

func (c ConfingMissing) Error() string {
	s := "Config missing: " + c.miss
	// fmt.Println(s)
	return s
}
