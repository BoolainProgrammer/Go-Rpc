package v1

import (
	"github.com/gin-gonic/gin"
	"toboefa/app/proto"
	"toboefa/app/transform/request"
	"toboefa/app/transform/response"
)

type ApiService struct{}

func (*ApiService) AddApiService(ctx *gin.Context) {
	req := new(request.ApiService)
	if err := ctx.ShouldBind(req); err != nil {
		response.HandleValidatorError(ctx, err)
		return
	}
	var resp proto.ApiServiceResp
	err := RpcClient.ApiServiceClient.AddApiService(proto.ApiServiceReq{ApiServiceName: req.ApiName}, &resp)
	if err != nil {
		return 
	}
}