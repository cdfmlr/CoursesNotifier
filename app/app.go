package app

import (
	"example.com/CoursesNotifier/courseTicker"
	"example.com/CoursesNotifier/jsonFileLoader"
	"example.com/CoursesNotifier/wx/wxAccessToken"
	"example.com/CoursesNotifier/wx/wxCoursesNotifier"
	"example.com/CoursesNotifier/wx/wxPlatformServer"
	"fmt"
	"net/http"
	"os"
	"time"
)

type App struct {
	conf    AppConf
	runtime AppRuntime
}

type AppConf struct {
	Wx       WxConf     `json:"wx"`
	Ticker   TickerConf `json:"ticker"`
	Database string     `json:"database"`
}

type WxConf struct {
	AppID                  string `json:"app_id"`
	AppSecret              string `json:"app_secret"`
	ReqToken               string `json:"req_token"`
	CourseNoticeTemplateID string `json:"course_notice_template_id"`
}

type TickerConf struct {
	timeToStart                string  `json:"time_to_start"`
	PeriodMinute               int     `json:"period_minute"`
	MinuteBeforeCourseToNotify float64 `json:"minute_before_course_to_notify"`
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

	// 初始化运行时组件
	a.initWxAccessTokenHolder()
	a.initWxPlatformServer()
	a.initCoursesTicker()

	return a
}

func (app *App) Run() {
	app.runWxPlatformServer()
	app.runCourseTicker()
}

// 初始化微信公众号 access_token holder
func (app *App) initWxAccessTokenHolder() {
	app.runtime.pWxAccessTokenHolder = wxAccessToken.NewHolder(app.conf.Wx.AppID, app.conf.Wx.AppSecret)
}

// 初始化微信公众号服务
func (app *App) initWxPlatformServer() {
	responser := wxPlatformServer.NewCourseNotifierResponser(app.conf.Database)
	app.runtime.pWxPlatformServer = wxPlatformServer.New(app.conf.Wx.ReqToken, responser, app.conf.Database)
}

// 初始化课程时钟
func (app *App) initCoursesTicker() {
	// 新建 课程时钟
	app.runtime.pCoursesTicker = courseTicker.NewCoursesTicker(
		"CourseTicker",
		app.conf.Database,
		time.Duration(app.conf.Ticker.PeriodMinute)*time.Minute,
		app.conf.Ticker.MinuteBeforeCourseToNotify,
		[]courseTicker.Notifier{
			courseTicker.LogNotifier("LogNotifier"),
			wxCoursesNotifier.New(
				app.conf.Wx.CourseNoticeTemplateID,
				app.runtime.pWxAccessTokenHolder,
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
	timeToStart, _ := time.Parse("2006-01-02 15:04", app.conf.Ticker.timeToStart)
	app.runtime.pCoursesTicker.Start(timeToStart)
}
