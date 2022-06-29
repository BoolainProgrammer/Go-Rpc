package test

import (
	"github.com/gin-gonic/gin"
	"sixstar/shopstar-micro/client/app/api"
	"sixstar/shopstar-micro/client/app/middleware"
)

var v1 = api.API.V1

func Tests(gin *gin.Engine)  {
	group := gin.Group("test")
	{
		// cache
		group.GET("/cache_get", v1.TestCache.GetCache)
		group.GET("/cache_set", v1.TestCache.SetCache)

		// jwt
		group.GET("/gen_jwt", v1.TestJwt.GenJwt)
		group.GET("/val_jwt",
			middleware.JwtAuth(),
			v1.TestJwt.ValJwt,
		)
		group.GET("/ref_jwt", v1.TestJwt.RefJwt)

		// validate ValidateSms
		group.POST("/validatesms",
			middleware.SmsValidateHandle(),
			v1.TestValidate.ValidateSms,
		)
		group.POST("/validatecaptcha",
			middleware.CaptchaValidateHandle(),
			v1.TestValidate.ValidateCaptcha,
		)

		group.POST("/validate",
			middleware.CaptchaValidateHandle(),
			middleware.SmsValidateHandle(),
			v1.TestValidate.Validate,
		)
	}
}
