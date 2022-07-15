package client

import (
	"toboefa/app/proto"
)

type ApiServiceClient struct {
}
//调用服务端添加apiservice接口
func (c *ApiServiceClient) AddApiService(req proto.ApiServiceReq, resp *proto.ApiServiceResp) error {
	return c.call("ApiServiceSever.AddApiService", req, resp)
}

//创建调度api_service服务
func (c *ApiServiceClient) call(severMethod string, req interface{}, resp interface{}) error {
	return call("api_service", severMethod, req, resp)
}
