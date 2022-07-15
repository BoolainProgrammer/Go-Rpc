package v1

import (
	"github.com/gin-gonic/gin"

)

type OrderCart struct {
}
// 添加订单
func (*OrderCart) AddCart(ctx *gin.Context)  {
	//var req request.AddCart
	//
	//if err := ctx.ShouldBind(&req); err != nil {
	//	response.HandleValidatorError(ctx, err)
	//	return
	//}
	//
	//// sku校验
	//
	//sku, err := goodsSkuService.GetSkuById(req.SkuId)
	//
	//if err != nil {
	//	response.FailWithMsg("无此商品", ctx)
	//	return
	//}
	//
	//if sku.Stock <= 0 {
	//	response.FailWithMsg("商品库存不足", ctx)
	//	return
	//}
	//
	//// 添加购物车
	//
	//authData := ctx.Value("auth_data").(util.AuthData)
	//cartId, err := orderCartService.AddCart(authData.MemberId(), req.SkuId, sku.GoodsId, req.Num, sku.Price, req.GoodsName, req.ShopName, req.Pricture)
	//
	//if err != nil {
	//	response.FailWithMsg("新增失败", ctx)
	//	return
	//}
	//
	//response.OkWithData(cartId, ctx)
}

func (*OrderCart) IndexCart(ctx *gin.Context)  {
	//response.OkWithData(orderCartService.GetCartByUserId(ctx.Query("user_id")), ctx)
}