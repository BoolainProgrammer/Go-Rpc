package api

import "sixstar/shopstar-micro/client/app/api/v1"

type Api struct {
	V1 v1.Api
}

var API = new(Api)
