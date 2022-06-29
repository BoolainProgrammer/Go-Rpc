package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"sixstar/shopstar-micro/client/global"
	"sixstar/shopstar-micro/client/app/util"
	"sixstar/shopstar-micro/client/app/transform/response"

)

func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("x-token")

		if token == "" {
			response.FailWithErr(global.ErrJwtAuthIllegal, ctx)
			return
		}

		if !util.JWT.Expired(token, global.JwtKey) {
			response.FailWithErr(global.ErrJwtAuthExpired, ctx)
			ctx.Done()
			return
		}

		data, err := util.JWT.ParseToken(token, global.JwtKey)
		if err != nil {
			response.FailWithErr(global.ErrJwtAuthFailure, ctx)
			return
		}
		fmt.Println(data)

		ctx.Set("auth_data", data)

		ctx.Next()
	}
}
