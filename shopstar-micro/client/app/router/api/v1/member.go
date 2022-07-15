package v1

import (
	"github.com/gin-gonic/gin"
	"sixstar/shopstar-micro/client/app/middleware"
)


func Member(gin *gin.Engine)  {
	group := gin.Group("member")
	{
		group.POST("/login",
			//middleware.CaptchaValidateHandle(),
			v1.Member.Login,
		)

		group.POST("/register",
			//middleware.CaptchaValidateHandle(),
			//middleware.SmsValidateHandle(),
			v1.Member.Register,
		)
		group.GET("/userinfo", middleware.JwtAuth(), v1.Member.UserInfo)

		group.POST("/regphonecheck", v1.Member.RegPhoneCheck)

		group.POST("/address", middleware.JwtAuth(), v1.MemberAddress.Create)
		group.GET("/defaultaddress", v1.MemberAddress.DefaultAddress)
	}
}
