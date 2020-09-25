# CoursesNotifier for NCEPU(Baoding)

> A courses Notifier for NCEPU(Baoding) based on [WeChat Official Accounts Platform](https://developers.weixin.qq.com/doc/offiaccount/en/Getting_Started/Overview.html).
>
> Designed for learning elementary Golang and SQL primer, this project is beginners friendly.

[微信公众号](https://developers.weixin.qq.com/doc/offiaccount/Getting_Started/Overview.html)上课时间提醒系统，为华北电力大学(保定)设计。

该项目为学习 Go 语言基础以及 SQL 入门而设计，适合初学者跟进练习。

## 功能实现

* 由用户提供学号、教务密码，系统自动从教务系统获取学生课表；
* 在每节课前通过微信公众号提醒订阅了该系统的学生；
* 期末出成绩自动通过微信公众号提醒订阅服务的学生。

## 项目基础

**开发技术**:

* Golang 基础
* SQL 基础
* 微信公众号服务开发基础

**基于项目**:

* 微信公众号服务： [腾讯/微信公众平台](https://developers.weixin.qq.com/doc/offiaccount/Getting_Started/Overview.html)
* 强智教务系统 API： [TLingC/QZAPI](https://github.com/TLingC/QZAPI/)
* 强智教务系统 Go SDK 实现：[cdfmlr/qzgo](https://github.com/cdfmlr/qzgo)

## 工作原理

由微信公众号服务组件与用户交互，用户可以通过公众号订阅、退订课程提醒业务。

当用户订阅课表时，提供学号以及教务密码，系统尝试依此登录学校的教务系统（强智教务系统），并尝试获取学生本学期课表，若成功则将信息写入数据库。

系统定时检查是否有课程快开始，若有则找出所有选修这些课的学生，通过微信公众号发送消息提醒上课。

同时，作为拓展，系统定时从教务系统获取当前学期课程成绩信息，如有更新，则推送通知学生。

## 拓展开发

本人系保定华电学生，在开发该系统时只考虑了本校实际情况。
但理论上任何使用**强智教务系统**的学校都可以通过简单的修改而使用该系统（修改学校名称、学号的数据库储存长度分配等即可）。

## 开发进度

- [x] 强智教务系统接入
- [x] 学生、课程、选课信息数据库
- [x] 微信公众号服务
- [x] 定时检测是否有课程快要开始
- [x] 微信课程提醒消息发送
- [x] BriefBullshitGenerator 增加系统趣味性
- [ ] 后台管理系统
- [ ] 安全性处理
- [ ] 消耗更小的持久化运行自动刷新、异常处理


## 欢迎贡献

欢迎任何人参与开发。贡献形式包括但不限于 Issues、Pull Requests、Forks。

## 预览

目前，我使用一个微信测试公众号部署了该服务。（由于开发未完成，不保障服务有效及信息安全）

欢迎关注，预览试用：

![微信测试公众号](https://tva1.sinaimg.cn/large/00831rSTgy1gcghbukj3gj30ci0a0dh1.jpg)

## 捐赠

鄙人穷苦，感谢支持🙏

* BTC: `1DgTSywmxeYvwpSxtNaU1zJE3VwK345v9A`
* LTC: `LXuQiCFc3JnzCd984WZmJ1MzFiJbBQobRa`

## 开源协议

Copyright 2020 CDFMLR

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
