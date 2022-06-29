package middleware

import (
	"time"

	"sixstar/shopstar-micro/client/global"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 请求记录是有用；可以做流量分析
func RequestLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 请求开始
		start := time.Now()
		// 请求路径
		path := ctx.Request.URL.Path
		// 执行请求
		ctx.Next()
		// 请求处理的结束时间
		end := time.Now()
		// 请求执行的时间
		cost := time.Since(start)
		// 获取状态
		status := ctx.Writer.Status()
		// 请求的基本信息

		// 可以自有扩展

		fields := []zap.Field{
			zap.Duration("start", start.Sub(start)),
			zap.Duration("end", start.Sub(end)),
			zap.Duration("cost", cost),
			// 基本信息
			zap.String("method", ctx.Request.Method),
			zap.Int("status", status),
			zap.String("errors", ctx.Errors.ByType(gin.ErrorTypePrivate).String()),
		}

		if status == 200 {
			global.Logs.Info(path, fields...)
		} else {
			global.Logs.Error(path, fields...)
		}
	}
}
