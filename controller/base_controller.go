package controller

import (
	"framework-go/common/function"
	"framework-go/pojo/request"
	"framework-go/services"
	"github.com/kataras/iris/v12"
	"github.com/qiafan666/gotato/commons"
	"github.com/qiafan666/gotato/commons/utils"
)

type BaseController struct {
	Ctx         iris.Context
	BaseService services.BaseService
}

func (receiver *BaseController) GetHealth() {
	receiver.Ctx.StatusCode(iris.StatusOK)
	return
}

func (receiver *BaseController) PostTest() {
	input := request.Test{}
	if code, msg := utils.ValidateAndBindCtxParameters(&input, receiver.Ctx, "BaseController PostTest"); code != commons.OK {
		_ = receiver.Ctx.JSON(commons.BuildFailedWithMsg(code, msg, input.RequestId))
		return
	}
	function.BindBaseRequest(&input, receiver.Ctx)
	if out, code, err := receiver.BaseService.Test(input); err != nil {
		_ = receiver.Ctx.JSON(commons.BuildFailed(code, input.Language, input.RequestId))
	} else {
		_ = receiver.Ctx.JSON(commons.BuildSuccess(out, input.Language, input.RequestId))
	}
}
