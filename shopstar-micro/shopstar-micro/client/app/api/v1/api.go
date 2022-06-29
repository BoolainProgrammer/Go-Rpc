package v1

import (
	tests "sixstar/shopstar-micro/client/app/api/v1/test"
	"sixstar/shopstar-micro/client/app/client"
)

//var (
//	PayService				 	 = service.Services.Pay
//
//	memberService 				 = service.Services.Member
//	memberAddressService		 = service.Services.MemberAddress
//
//	SysAreaService		 		 = service.Services.SysArea
//
//	goodsService 	 			 = service.Services.Goods
//	goodsCategoryService		 = service.Services.GoodsCategory
//	goodsSkuService   	 		 = service.Services.GoodsSku
//
//	orderService				 = service.Services.Order
//	orderCartService			 = service.Services.OrderCart
//)

var RpcClient = client.RpcClient

type Api struct {
	Order
	OrderCart

	Pay

	Member
	MemberAddress

	Goods
	GoodsSku
	GoodsCategory

	Verification

	tests.TestJwt
	tests.TestCache
	tests.TestValidate
}
