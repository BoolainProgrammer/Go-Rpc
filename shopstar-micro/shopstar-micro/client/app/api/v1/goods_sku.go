package v1

import (
	"github.com/gin-gonic/gin"

)

type GoodsSku struct {
	
}
// 商品的sku列表
func (*GoodsSku)  AssemblyAttr (ctx *gin.Context) {
	//goodsId,_ := strconv.Atoi(ctx.Query("goods_id"))
	//
	//if goodsId == 0 {
	//	response.FailWithMsg("无此商品", ctx)
	//	return
	//}
	//
	//attrs := goodsSkuService.AssemblyAttr(int64(goodsId))
	//
	//response.OkWithData(response.CateSku{
	//	Sku: attrs,
	//}, ctx)
}
// 商品筛选
func (*GoodsSku) Tbdetail(ctx *gin.Context) {
	//// 单纯的查询列表
	//req := new(request.GoodsSku)
	//
	//if err := ctx.ShouldBind(req); err != nil {
	//	response.HandleValidatorError(ctx, err)
	//	return
	//}
	//
	//skulist,err := goodsSkuService.SkuAttrCollection(req.GoodsId, req.SkuProperties)
	//if err != nil {
	//	response.FailWithErr(err, ctx)
	//	return
	//}
	//
	//response.OkWithData(response.Tbdetail{
	//	Data: skulist,
	//}, ctx)
}

