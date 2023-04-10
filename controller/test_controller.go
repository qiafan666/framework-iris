package controller

import (
	"framework-go/common/function"
	"framework-go/pojo/request"
	"framework-go/services"
	"github.com/kataras/iris/v12"
	"github.com/qiafan666/quickweb/commons"
	"github.com/qiafan666/quickweb/commons/utils"
)

type TestController struct {
	Ctx         iris.Context
	TestService services.TestService
}

func (receiver *TestController) PostTest() {
	input := request.Test{}
	if code, msg := utils.ValidateAndBindCtxParameters(&input, receiver.Ctx, "PortalWebController PostRegister"); code != commons.OK {
		_ = receiver.Ctx.JSON(commons.BuildFailedWithMsg(code, msg))
		return
	}
	function.BindBaseRequest(&input, receiver.Ctx)
	if out, code, err := receiver.TestService.Test(input); err != nil {
		_ = receiver.Ctx.JSON(commons.BuildFailed(code, input.Language))
	} else {
		_ = receiver.Ctx.JSON(commons.BuildSuccess(out, input.Language))
	}
}
