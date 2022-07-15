package service

import (
	"toboefa/app/logic"
	"toboefa/app/proto"
)

type ApiServiceServer struct {

}

func NewApiServiceServer() *ApiServiceServer {
	return &ApiServiceServer{}
}

func (*ApiServiceServer)AddApiService(req proto.AddApiServiceReq) error {
	return logic.ApiServiceLogic.AddApiService(req)

}