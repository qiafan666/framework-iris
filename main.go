package main

import (
	"framework-go/common"
	router "framework-go/controller"
	"github.com/qiafan666/gotato"
	"github.com/qiafan666/gotato/commons"
)

// @title framework API Document
// @description framework API Document
// @version 1
// @schemes http
// @produce json
// @consumes json
func main() {
	server := gotato.GetGotatoInstance()
	server.Default()
	server.RegisterErrorCodeAndMsg(commons.MsgLanguageEnglish, common.EnglishCodeMsg)
	//server.StartServer(cornus.DatabaseService, cornus.OssService)
	router.RegisterRouter(server.App().GetIrisApp())
	server.WaitClose()
}
