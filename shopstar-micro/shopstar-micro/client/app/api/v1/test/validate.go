package tests

import "github.com/gin-gonic/gin"

type TestValidate struct {

}
/**
 * @Com www.sixstaredu.com
 * @Author 六星教育-shineyork老师
 */
func (*TestValidate) ValidateSms(ctx *gin.Context) {
	ctx.String(200, "okok-ValidateSms" )
}
func (*TestValidate) ValidateCaptcha(ctx *gin.Context) {
	ctx.String(200, "okok-ValidateCaptcha" )
}
func (*TestValidate) Validate(ctx *gin.Context) {
	ctx.String(200, "okok-Validate" )
}