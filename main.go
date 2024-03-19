package main

import (
	"framework-iris/common"
	router "framework-iris/controller"
	"framework-iris/gota"
	"framework-iris/gota/commons"
)

// @title framework API Document
// @description framework API Document
// @version 1
// @schemes http
// @produce json
// @consumes json
func main() {
	server := gota.GetgotaInstance()
	server.Default()
	server.RegisterErrorCodeAndMsg(commons.MsgLanguageEnglish, common.EnglishCodeMsg)
	//server.StartServer(cornus.DatabaseService, cornus.OssService)
	router.RegisterRouter(server.App().GetIrisApp())
	server.WaitClose()
}
