package v1

import (
	"github.com/gin-gonic/gin"

	"sixstar/shopstar-micro/client/app/proto"
	"sixstar/shopstar-micro/client/app/transform/request"
	"sixstar/shopstar-micro/client/app/transform/response"
	"sixstar/shopstar-micro/client/app/util"
	"sixstar/shopstar-micro/client/global"
)

type Member struct {

}

func (*Member) Login(ctx *gin.Context) {
	// 参数匹配
	req := new(request.Login)
	if err := ctx.ShouldBind(req); err != nil {
		response.HandleValidatorError(ctx, err)
		return
	}
	//// 登入-验证
	var resp proto.MemberResp
	err := RpcClient.MemberClient.Login(proto.LoginReq{
		Username: req.Username,
		Password: req.Password,
	}, &resp)
	if err != nil {
		response.FailWithMsg(err.Error(), ctx)
		return
	}
	//model,err := memberService.Login(req.Username, req.Password)
	//if err != nil {
	//	response.FailWithMsg(err.Error(), ctx)
	//	return
	//}
	// token生成
	token, err := util.JWT.GenerateToken(global.JwtKey, util.NewAuthData(resp.Id))

	if err != nil  {
		response.FailWithMsg("登入失败", ctx)
		return
	}

	response.OkWithData(gin.H{"token":token}, ctx)
}

func (*Member) UserInfo(ctx *gin.Context) {
	authData := ctx.Value("auth_data").(util.AuthData)

	var resp proto.MemberResp
	err := RpcClient.MemberClient.UserInfo(proto.MemberReq{
		Id: authData.MemberId(),
	}, &resp)
	if err != nil {
		response.FailWithErr(err, ctx)
		return
	}

	response.OkWithDetailed(resp, "查询成功", ctx)
}
// 如果是已有1000w用户，然后快速处理登入验证问题
func (*Member) Register(ctx *gin.Context)  {
	//req := new(request.Register)
	//
	//if err := ctx.ShouldBind(req); err != nil {
	//	response.HandleValidatorError(ctx, err)
	//	return
	//}
	//
	//_, err := memberService.Register(req.Username, req.Password, req.Phone)
	//
	//if err != nil {
	//	response.FailWithMsg(err.Error(), ctx)
	//	return
	//}
	//// 17680143620
	//response.OkWithMsg("创建成功", ctx)
}
// 验证手机号码是否登入过
func (*Member) RegPhoneCheck(ctx *gin.Context) {
	//phone := ctx.PostForm("phone")
	//if phone == "" {
	//	response.OkWithMsg("暂未存在", ctx)
	//	return
	//}
	//
	//if _, err := memberService.FindMemberByPhone(phone); err == nil {
	//	response.OkWithMsg("暂未存在", ctx)
	//	return
	//}
	//
	//response.FailWithMsg("存在手机号码", ctx)
}
