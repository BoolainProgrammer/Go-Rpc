package v1

import (
	"github.com/gin-gonic/gin"
)

func Pay(gin *gin.Engine)  {

	gin.POST("/alipay", v1.Pay.AliPay)
	gin.POST("/alinotifyUrl", v1.Pay.NotifyUrl)
	// ConfirmOrder
}

