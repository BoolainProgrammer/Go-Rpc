package middleware

import (
	"github.com/gin-gonic/gin"
	"sixstar/shopstar-micro/client/core/cache"
	"sixstar/shopstar-micro/client/global"
	"sixstar/shopstar-micro/client/app/transform/response"
)


//
//  SmsValidateHandle
//  @Description: 短信验证统一处理
//  @return gin.HandlerFunc
//
func SmsValidateHandle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key := ctx.PostForm("verification_key")
		code, err := cache.CacheManager.Get(key)

		if err != nil || string(code.([]byte)) != ctx.PostForm("sms_code") {
			response.FailWithErr(global.ErrValidatorSmsCode, ctx)
			ctx.Abort()
			return
		}

		cache.CacheManager.Delete(key)
		// 正常的逻辑
		ctx.Next()
	}
}
//
//  CaptchaValidateHandle
//  @Description: 验证码统一验证处理方法
//  @return gin.HandlerFunc
//
func CaptchaValidateHandle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		captchaid := ctx.PostForm("captchaid")
		captcha_code := ctx.PostForm("captcha_code")

		if captchaid == "" || captcha_code == "" {
			response.FailWithErr(global.ErrValidatorCaptahCode, ctx)
			ctx.Abort()
			return
		}

		if !global.CaptchaStore.Verify(captchaid, captcha_code, true) {
			response.FailWithErr(global.ErrValidatorCaptahCode, ctx)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
