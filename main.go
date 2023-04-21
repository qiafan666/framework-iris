package main

import (
	"framework-go/common"
	router "framework-go/controller"
	"github.com/qiafan666/gotato"
	"github.com/qiafan666/gotato/commons"
)

func main() {
	server := gotato.GetGotatoInstance()
	server.Default()
	server.RegisterErrorCodeAndMsg(commons.MsgLanguageEnglish, common.EnglishCodeMsg)
	//server.StartServer(cornus.DatabaseService, cornus.OssService)
	router.RegisterRouter(server.App().GetIrisApp())
	server.WaitClose()
}
