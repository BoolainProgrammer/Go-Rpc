package tests

import (
	"github.com/gin-gonic/gin"
	"sixstar/shopstar-micro/client/global"
	"sixstar/shopstar-micro/client/app/util"
)

type TestJwt struct {

}
/**
 * @Com www.sixstaredu.com
 * @Author 六星教育-shineyork老师
 */
func (*TestJwt) GenJwt(ctx *gin.Context) {
	token, _ := util.JWT.GenerateToken(global.JwtKey, map[string]interface{}{
		"member_id": "0x000008a000000001",
	})
	ctx.String(200, token)
}
func (*TestJwt) ValJwt(ctx *gin.Context) {
	ctx.String(200, "ok - ValJwt")
}

func (*TestJwt) RefJwt(ctx *gin.Context) {
	token,err := util.JWT.RefreshToken(ctx.Request.Header.Get("x-token"), global.JwtKey)
	if err != nil {
		ctx.String(200, err.Error())
		return
	}
	ctx.String(200, token)
}