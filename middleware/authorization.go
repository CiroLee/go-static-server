package middleware

import (
	"github.com/CiroLee/go-static-server/config"
	"github.com/CiroLee/go-static-server/response"
	"github.com/gin-gonic/gin"
)

func Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if cookie, err := ctx.Request.Cookie("token"); err == nil {
			token := cookie.Value
			// 有权限，继续
			if token == config.Token {
				ctx.Next()
				return
			}
		}
		response.UnAuth(ctx)
		ctx.Abort()
	}
}
