package v1

import (
	"github.com/gin-gonic/gin"
	"sixstar/shopstar-micro/client/app/middleware"
)

func Order(gin *gin.Engine)  {

	group := gin.Group("order", middleware.JwtAuth())
	{
		group.GET("/cart", v1.OrderCart.IndexCart)
		group.POST("/cart", v1.OrderCart.AddCart)

		group.POST("/create", v1.Order.Create)
	}

	// ConfirmOrder
}

