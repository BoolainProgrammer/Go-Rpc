package request
/**
 * @Com www.sixstaredu.com
 * @Author 六星教育-shineyork老师
 */

type GoodsSku struct {
	GoodsId int64 		`form:"goods_id" binding:"required"`
	SkuProperties string `form:"sku_properties" binding:"required"`
}