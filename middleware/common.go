package middleware

import (
	"context"
	"framework-go/common"
	"framework-go/pojo/request"
	"github.com/kataras/iris/v12"
	"github.com/qiafan666/gotato/commons"
	"io"
	"net/http"
)

func Common(ctx iris.Context) {
	//get language
	language := ctx.Request().Header.Get("Language")
	if language == "" {
		language = commons.DefaultLanguage
	}
	ctx.Values().Set(common.BaseRequest, request.BaseRequest{
		Ctx:      ctx.Values().Get("ctx").(context.Context),
		Language: language,
	})
	if ctx.Request().Method == http.MethodPost {

		body, err := io.ReadAll(ctx.Request().Body)
		if err != nil {
			_ = ctx.JSON(commons.BuildFailed(commons.UnKnowError, commons.DefaultLanguage))
			return
		}

		ctx.Values().Set(commons.CtxValueParameter, body)
	}
	ctx.Next()
}
