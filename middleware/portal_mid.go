package middleware

import (
	"framework-iris/common"
	cornus "framework-iris/gota"
	"framework-iris/gota/commons"
	"framework-iris/pojo/request"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/v12"
)

var jwtConfig struct {
	JWT struct {
		Secret string `yaml:"secret"`
	} `yaml:"jwt"`
}

func init() {
	cornus.GetgotaInstance().LoadCustomizeConfig(&jwtConfig)
}

var witheList = map[string]string{
	"/v1/test":   "",
	"/v1/health": "",
}

func CheckPortalAuth(ctx iris.Context) {

	var language, requestId, phone, companyName string
	var userId, roleId int

	BaseRequest := ctx.Values().Get(common.BaseRequest).(request.BaseRequest)
	language = BaseRequest.Language
	requestId = BaseRequest.RequestId
	//check white list
	if _, ok := witheList[ctx.Request().RequestURI]; !ok {

		//check jwt
		parseToken, err := jwt.Parse(ctx.Request().Header.Get("Authorization"), func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtConfig.JWT.Secret), nil
		})
		if err != nil {
			_ = ctx.JSON(commons.BuildFailed(commons.TokenError, language, requestId))
			return
		}

		if _, ok := parseToken.Claims.(jwt.MapClaims); ok && parseToken.Valid {

		} else {
			_ = ctx.JSON(commons.BuildFailed(commons.UnKnowError, language, requestId))
			return
		}

	}

	ctx.Values().Set(common.BasePortalRequest, request.BasePortalRequest{
		BaseID:      int64(userId),
		Phone:       phone,
		Role:        roleId,
		CompanyName: companyName,
	})
	ctx.Next()
}
