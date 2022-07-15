package v1

import (
	"sixstar/shopstar-micro/client/app/middleware"
	"sixstar/shopstar-micro/client/app/router/api/v1/test"
	"sixstar/shopstar-micro/client/core/router"
	"sixstar/shopstar-micro/client/app/api"

)

var v1 = api.API.V1

func InitRouter()  {
	router.Use(
		middleware.Recover(),
		middleware.RequestLog(),
	)

	router.Register(
		Member,
		Goods,
		Order,
		Pay,
		Verification,

		test.Tests,
	)
}
