package v1

import (
	"github.com/gin-gonic/gin"
	//"sixstar/shopstar-micro/client/app/transform/response"
	//"strconv"
)

type GoodsCategory struct {
}

func (*GoodsCategory) GetGoodsCategoryByParentId (ctx *gin.Context) {
	//parentId := ctx.PostForm("parent_id")
	//if parentId == "" {
	//	response.Fail(ctx)
	//	return
	//}
	//
	//pid,_ := strconv.Atoi(parentId)
	//
	//data,err := goodsCategoryService.GetCategoryListByParentId(pid)
	//
	//if err != nil {
	//	response.FailWithMsg("读取错误", ctx)
	//	return
	//}
	//
	//response.OkWithData(data, ctx)
}