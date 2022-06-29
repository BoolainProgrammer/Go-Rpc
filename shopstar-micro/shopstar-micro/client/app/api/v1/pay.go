package v1

import (
	"github.com/gin-gonic/gin"
	//"sixstar/shopstar-micro/client/app/transform/request"
	//"sixstar/shopstar-micro/client/app/transform/response"

)

type Pay struct {
	
}

func (*Pay) AliPay(ctx *gin.Context) {
	//var rep request.Pay
	//
	//if err := ctx.ShouldBind(&rep); err != nil {
	//	return
	//}
	//
	//url, err := PayService.AliPay(rep.OrderId)
	//if err != nil {
	//	return
	//}
	//
	//response.OkWithData(response.Pay{
	//	PayQrcodeUrl: url,
	//}, ctx)
}

func (*Pay) NotifyUrl(ctx *gin.Context)  {
	//var rep request.AliUnifyPay
	//
	//if err := ctx.ShouldBind(&rep); err != nil {
	//	response.HandleValidatorError(ctx, err)
	//	return
	//}
	//
	//orderGoods,err := PayService.AliNotifyPay(ctx.Request)
	//
	//ctx.String(200, "Success")
	//
	//ctx.Set("orderGoods", orderGoods)
	//
	//if err == nil {
	//	fmt.Println("alipay success !!!")
	//}
}