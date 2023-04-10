package main

import (
	"framework-go/common"
	router "framework-go/controller"
	cornus "github.com/qiafan666/quickweb"
	"github.com/qiafan666/quickweb/commons"
)

func main() {
	server := cornus.GetCornusInstance()
	server.Default()
	server.RegisterErrorCodeAndMsg(commons.MsgLanguageEnglish, common.EnglishCodeMsg)
	server.StartServer(cornus.DatabaseService, cornus.OssService)
	router.RegisterRouter(server.App().GetIrisApp())
	server.WaitClose()
}
