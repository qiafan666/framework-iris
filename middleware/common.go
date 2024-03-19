package middleware

import (
	"context"
	"framework-iris/common"
	"framework-iris/gota/commons"
	"framework-iris/gota/middleware"
	"framework-iris/pojo/request"
	"github.com/kataras/iris/v12"
	"io"
	"net/http"
	"sync"
)

var blackList = []string{
	"/v1/health",
	"/swagger/*",
	"/debug/pprof/*",
}

var once sync.Once

func init() {
	once.Do(func() {
		middleware.RegisterIgnoreRequest(blackList...)
	})
}

func Common(ctx iris.Context) {

	//get language
	language := ctx.Request().Header.Get("Language")
	if language == "" {
		language = commons.DefaultLanguage
	}
	c := ctx.Values().Get("ctx").(context.Context)
	requestId := c.Value("trace_id").(string)
	ctx.Values().Set(common.BaseRequest, request.BaseRequest{
		Ctx:       c,
		RequestId: requestId,
		Language:  language,
	})
	if ctx.Request().Method == http.MethodPost {

		body, err := io.ReadAll(ctx.Request().Body)
		if err != nil {
			_ = ctx.JSON(commons.BuildFailed(commons.UnKnowError, commons.DefaultLanguage, requestId))
			return
		}

		ctx.Values().Set(commons.CtxValueParameter, body)
	}

	ctx.Next()
}
