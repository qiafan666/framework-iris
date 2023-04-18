package controller

import (
	"framework-go/middleware"
	"framework-go/services"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func RegisterRouter(ctx *iris.Application) {
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, //允许通过的主机名称
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	})
	ctx.Use(middleware.Common)
	//web router
	mvc.Configure(ctx.Party("/v1", crs).AllowMethods(iris.MethodOptions, iris.MethodPut),
		func(application *mvc.Application) {
			application.Router.Use(middleware.CheckPortalAuth)
			application.Handle(&BaseController{
				BaseService: services.NewBaseServiceInstance(),
			})
		})
}
