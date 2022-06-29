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
		group.POST("/save-api",
			v1.ApiService.SaveApiService,
		)
		group.GET("/find-api-list",
			v1.ApiService.FindApiServiceList,
		)
		group.GET("/find-api-by-id",
			v1.ApiService.FindApiServiceById,
		)

	}
}
