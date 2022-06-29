package api

import "toboefa/app/api/v1"

type Api struct {
	V1 v1.Api
}

var API = new(Api)
