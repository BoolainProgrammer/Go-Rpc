package request

type OrderCreate struct {
	// 购物车列表
	// sku:num
	SkuId int `form:"sku_id" json:"sku_id" binding:"required"`
	SkuNum int `form:"sku_num" json:"sku_num" binding:"required"`
	//GoodsSkuList []string `form:"goods_sku_list" json:"goods_sku_list" binding:"required"`
	// 优惠券
	UseCoupon int `form:"use_coupon" json:"use_coupon"`
	// 积分
	Integral int `form:"integral" json:"integral"`
	// 留言
	LeaveMsg string `form:"leave_msg" json:"leave_msg"`
	// 收货地址 根据用户获取收货地址信息
	ReceivingAddress int `form:"receiving_address" json:"receiving_address" binding:"required"`
	// 配送服务 ; 根据商家的特点来，免运费或需要运费
	Distribution byte `form:"distribution" json:"distribution" binding:"required"`
}
// userId 用户id
// skuList sku列表
// useCoupon 用户优惠券
// receivingAddress 收货地址
// distribution 配送方式
// leaveMsg 备注信息
// source 来源