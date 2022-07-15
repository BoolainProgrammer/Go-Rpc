package v1

import (
	"github.com/gin-gonic/gin"
)

func apiService(gin *gin.Engine) {
	group := gin.Group("api")
	{
		group.POST("/add-api",
			v1.ApiService.AddApiService,
		)


	}
}
