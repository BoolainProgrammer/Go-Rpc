package request

/**
 * @Com www.sixstaredu.com
 * @Author 六星教育-shineyork老师
 */
// userId string, skuId, goodsId, num int, price decimal.Decimal, goodsName, shopName, picture string
type AddCart struct {
	SkuId  int  `form:"sku_id"  binding:"required"`
	Num    int    `form:"num"  binding:"required,gte=1"`
	GoodsName string `form:"goods_name" binding:"required"`
	ShopName string `form:"shop_name" binding:"required"`
	Pricture string `form:"pricture"`
}
