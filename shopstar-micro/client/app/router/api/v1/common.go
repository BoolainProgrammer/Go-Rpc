package v1

import "github.com/gin-gonic/gin"

func Verification(gin *gin.Engine)  {
	gin.POST("/smscode", v1.Verification.SmsCodes)
	gin.GET("/captcha", v1.Verification.Captcha)
}
