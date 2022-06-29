package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"runtime/debug"

	"sixstar/shopstar-micro/client/global"
	"sixstar/shopstar-micro/client/app/transform/response"
)

//
//  Recover
//  @Description: 用于对系统运行中产生的致命异常进行记录
//  @return gin.HandlerFunc
//
func Recover() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// panic, recover

		defer func() {
			if err := recover(); err != nil {
				// 记录
				global.Logs.Error(
					ctx.Request.URL.Path,
					zap.Any("error", err),
					zap.String("stack", string(debug.Stack())),
				)
				fmt.Println(err)
				response.FailWithMsg("系统错误", ctx)
				return
			}
		}()

		ctx.Next()

	}
}
