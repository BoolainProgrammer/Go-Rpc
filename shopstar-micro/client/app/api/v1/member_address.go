package v1

import (
	"github.com/gin-gonic/gin"
	//"sixstar/shopstar-micro/client/app/transform/request"
	//"sixstar/shopstar-micro/client/app/transform/response"
	//"sixstar/shopstar-micro/client/app/util"
)

type MemberAddress struct {

}
// 默认地址
func (*MemberAddress) DefaultAddress(ctx *gin.Context)  {
	//id := ctx.Query("user_id")
	//address,_ := memberAddressService.GetMemberAddressInfo(id)
	//
	//province, city, district := SysAreaService.GetAddress(address.Province, address.City, address.District)
	//
	//response.OkWithData(&response.MemberAddress{
	//	Address: address,
	//	Province: province.ProvinceName,
	//	City: city.CityName,
	//	District: district.DistrictName,
	//}, ctx)
}
// 新增收货地址
func (*MemberAddress) Create(ctx *gin.Context) {
	//req := new(request.MemberAddress)
	//if err := ctx.ShouldBind(req); err != nil {
	//	response.HandleValidatorError(ctx, err)
	//	return
	//}
	//
	//authData := ctx.Value("auth_data").(util.AuthData)
	//
	//err := memberAddressService.AddMemberAddress(authData.MemberId(), req.Recipient, req.Phone, req.DistrictId, req.CityId, req.DistrictId, req.Address)
	//if err != nil {
	//	response.FailWithMsg("新增失败", ctx)
	//	return
	//}
	//
	//response.OkWithMsg("新增成功", ctx)
}