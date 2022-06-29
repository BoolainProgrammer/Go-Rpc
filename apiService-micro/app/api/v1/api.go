package v1

import (
	"toboefa/app/service"
)

var (
	ApiServices 	 			 = service.Services.Api
	MenuServices 	 			 = service.Services.Menu
)

type Api struct {
	ApiService
	Menu
}
