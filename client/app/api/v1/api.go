package v1

import (
	"toboefa/app/client"
)

type Api struct {
	ApiService
}

var RpcClient = client.RpaClient
