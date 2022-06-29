package middleware

import (

	"github.com/gin-gonic/gin"

	"sixstar/shopstar-micro/member-server/global"
	"sixstar/shopstar-micro/member-server/app/util"
	"sixstar/shopstar-micro/member-server/app/transform/response"

)


//
//  JwtAuth
//  @Description: 用于jwt权限验证
//  @return gin.HandlerFunc
//
func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("x-token")

		if token == "" {
			//response.FailWithErr(global.ErrJwtAuthIllegal, ctx)
			return
		}

		if !util.JWT.Expired(token, global.JwtKey) {
			//response.FailWithErr(global.ErrJwtAuthExpired, ctx)
			ctx.Done()
			return
		}

		data, err := util.JWT.ParseToken(token, global.JwtKey)
		if err != nil {
			//response.FailWithErr(global.ErrJwtAuthFailure, ctx)
			return
		}

		ctx.Set("auth_data", data)

		ctx.Next()
	}
}
