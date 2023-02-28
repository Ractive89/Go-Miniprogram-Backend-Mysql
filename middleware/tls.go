package middleware

import (
	"dayang/conf"

	"github.com/unrolled/secure"

	"github.com/gin-gonic/gin"
)

func TlsHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		middleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     conf.HostName + conf.HttpPort,
		})
		err := middleware.Process(ctx.Writer, ctx.Request)

		//如果出现错误，请不要继续。
		if err != nil {
			ctx.Abort()
			return
		}

		if status := ctx.Writer.Status(); status > 300 && status < 399 {
			ctx.Abort()
		}

		ctx.Next()
	}
}
