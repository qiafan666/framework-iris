package controller

import (
	"framework-iris/common/function"
	"framework-iris/pojo/request"
	"framework-iris/services"
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

// PostTest godoc
// @Summary Test
// @Description Test
// @Tags test
// @Accept  json
// @Produce  json
// @Router /v1/test [post]
// @param data body request.Test true "request.Test"
// @Success 200 {object} response.Test
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
