package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		clientIP := ctx.ClientIP()
		userAgent := ctx.Request.UserAgent()
		method := ctx.Request.Method
		statusCode := ctx.Writer.Status()
		log.Printf("%15s | %13v | %s | %s | %s\n",
			clientIP,
			statusCode,
			userAgent,
			method,
			ctx.Request.URL.Path,
		)
	}
}
