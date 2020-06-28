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
	"example.com/CoursesNotifier/examResultTicker"
	"example.com/CoursesNotifier/util/jsonFileLoader"
	"example.com/CoursesNotifier/wx/wxAccessToken"
	"example.com/CoursesNotifier/wx/wxCoursesNotifier"
	"example.com/CoursesNotifier/wx/wxExamResultNotifier"
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
	Wx               WxConf               `json:"wx"`
	CourseTicker     CourseTickerConf     `json:"course_ticker"`
	ExamResultTicker ExamResultTickerConf `json:"exam_result_ticker"`
	Data             DataConf             `json:"data"`
}

type WxConf struct {
	AppID                      string `json:"app_id"`
	AppSecret                  string `json:"app_secret"`
	ReqToken                   string `json:"req_token"`
	CourseNoticeTemplateID     string `json:"course_notice_template_id"`
	ExamResultNoticeTemplateID string `json:"exam_result_notice_template_id"`
}

type CourseTickerConf struct {
	TimeToStart                string  `json:"time_to_start"`
	PeriodMinute               int     `json:"period_minute"`
	MinuteBeforeCourseToNotify float64 `json:"minute_before_course_to_notify"`
}

type ExamResultTickerConf struct {
	TimeToStart  string `json:"time_to_start"`
	PeriodMinute int    `json:"period_minute"`
}

type DataConf struct {
	Database         string `json:"database"`
	BullshitDataFile string `json:"bullshit_data_file"`
}

type AppRuntime struct {
	pWxAccessTokenHolder *wxAccessToken.Holder
	pCoursesTicker       *courseTicker.CoursesTicker
	pExamResultTicker    *examResultTicker.ExamResultTicker
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
	errs := make([]*ConfigMissing, 0)

	// CourseTicker
	if app.conf.CourseTicker.PeriodMinute == int(0) {
		errs = append(errs, NewConfigMissing("CourseTicker.PeriodMinute"))
	}
	if strings.TrimSpace(app.conf.CourseTicker.TimeToStart) == "" {
		errs = append(errs, NewConfigMissing("CourseTicker.TimeToStart"))
	}
	if app.conf.CourseTicker.MinuteBeforeCourseToNotify == float64(0) {
		errs = append(errs, NewConfigMissing("CourseTicker.MinuteBeforeCourseToNotify"))
	}

	// ExamResultTicker
	if app.conf.ExamResultTicker.PeriodMinute == 0 {
		errs = append(errs, NewConfigMissing("ExamResultTicker.PeriodMinute"))
	}
	if strings.TrimSpace(app.conf.ExamResultTicker.TimeToStart) == "" {
		errs = append(errs, NewConfigMissing("ExamResultTicker.TimeToStart"))
	}

	// Wx
	if strings.TrimSpace(app.conf.Wx.AppID) == "" {
		errs = append(errs, NewConfigMissing("Wx.AppID"))
	}
	if strings.TrimSpace(app.conf.Wx.AppSecret) == "" {
		errs = append(errs, NewConfigMissing("Wx.AppSecret"))
	}
	if strings.TrimSpace(app.conf.Wx.CourseNoticeTemplateID) == "" {
		errs = append(errs, NewConfigMissing("Wx.CourseNoticeTemplateID"))
	}
	if strings.TrimSpace(app.conf.Wx.ExamResultNoticeTemplateID) == "" {
		errs = append(errs, NewConfigMissing("Wx.ExamResultNoticeTemplateID"))
	}
	if strings.TrimSpace(app.conf.Wx.ReqToken) == "" {
		errs = append(errs, NewConfigMissing("Wx.ReqToken"))
	}

	// Data
	if strings.TrimSpace(app.conf.Data.Database) == "" {
		errs = append(errs, NewConfigMissing("Data.Database"))
	}
	if strings.TrimSpace(app.conf.Data.BullshitDataFile) == "" {
		errs = append(errs, NewConfigMissing("Data.BullshitDataFile"))
	}

	if len(errs) != 0 {
		s := ""
		for _, e := range errs {
			s += e.miss + ", "
		}
		return *NewConfigMissing(strings.Trim(s, ", "))
	}
	return nil
}

func (app *App) Run() {
	// 初始化运行时组件
	app.initWxAccessTokenHolder()
	app.initWxPlatformServer()
	app.initCoursesTicker()
	app.initExamResultTicker()

	// 启动守护任务
	app.runWxPlatformServer()
	app.runCourseTicker()
	app.runExamResultTicker()
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
		time.Duration(app.conf.CourseTicker.PeriodMinute)*time.Minute,
		app.conf.CourseTicker.MinuteBeforeCourseToNotify,
		[]courseTicker.Notifier{
			courseTicker.LogNotifier("LogNotifier >> course"),
			wxCoursesNotifier.New(
				app.conf.Wx.CourseNoticeTemplateID,
				app.runtime.pWxAccessTokenHolder,
				app.conf.Data.BullshitDataFile,
			),
		},
	)
}

// 初始化考试成绩时钟
func (app *App) initExamResultTicker() {
	app.runtime.pExamResultTicker = examResultTicker.NewExamResultTicker(
		"ExamResultTicker",
		app.conf.Data.Database,
		time.Duration(app.conf.ExamResultTicker.PeriodMinute)*time.Minute,
		[]examResultTicker.Notifier{
			examResultTicker.LogNotifier("LogNotifier >> examResult"),
			wxExamResultNotifier.New(
				app.conf.Wx.ExamResultNoticeTemplateID,
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
	timeToStart, _ := time.Parse("2006-01-02 15:04", app.conf.CourseTicker.TimeToStart)
	app.runtime.pCoursesTicker.Start(timeToStart)
}

// 开始运行考试成绩时钟
func (app *App) runExamResultTicker() {
	timeToStart, _ := time.Parse("2006-01-02 15:04", app.conf.ExamResultTicker.TimeToStart)
	app.runtime.pExamResultTicker.Start(timeToStart)
}

type ConfigMissing struct {
	miss string
}

func NewConfigMissing(miss string) *ConfigMissing {
	return &ConfigMissing{miss: miss}
}

func (c ConfigMissing) Error() string {
	s := "Config missing: " + c.miss
	// fmt.Println(s)
	return s
}
