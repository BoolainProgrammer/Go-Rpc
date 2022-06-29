package v1

import (
	"toboefa/app/api"
	"toboefa/app/middleware"
	"toboefa/core/router"
)

var v1 = api.API.V1

func InitRouter()  {
	router.Use(
		middleware.Recover(),
		middleware.RequestLog(),
	)

	router.Register(
		apiService,
		menu,
	)
}
