package v1

import (
	"github.com/gin-gonic/gin"
)

type Goods struct {}

func (*Goods) AttributeList(ctx *gin.Context)  {
	//goodsId,_ := strconv.Atoi(ctx.Query("goods_id"))
	//
	//attr,_ := goodsService.GetGoodsAttributeList(int64(goodsId))
	//
	//response.OkWithData(response.AttributeList{
	//	Attrs: attr,
	//}, ctx)
}


