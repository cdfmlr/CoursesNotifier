package wxPlatformServer

type Responser interface {
	Do(reqUser string, reqContent string) (respContent string)
}
