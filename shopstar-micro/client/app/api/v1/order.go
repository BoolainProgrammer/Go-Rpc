package v1

import (
	"github.com/gin-gonic/gin"
	//"sixstar/shopstar-micro/client/app/util"
	//
	//"sixstar/shopstar-micro/client/app/transform/request"
	//"sixstar/shopstar-micro/client/app/transform/response"
)

type Order struct {
	
}

func (*Order) OrderInfo(ctx *gin.Context) {
	
}
// 创建订单
func (*Order) Create(ctx *gin.Context) {
	//var req request.OrderCreate
	//
	//if err := ctx.BindJSON(&req); err != nil {
	//	response.HandleValidatorError(ctx, err)
	//	return
	//}
	//authData := ctx.Value("auth_data").(util.AuthData)
	//
	//orderNo, err := orderService.CreateOrder(authData.MemberId(), req.SkuId, req.SkuNum, req.UseCoupon, req.ReceivingAddress, req.Distribution, req.LeaveMsg)
	//
	//if err != nil {
	//	response.FailWithErr(err, ctx)
	//	return
	//}
	//
	//response.OkWithData(orderNo, ctx)
}
