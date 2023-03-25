package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func InterceptorMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Method != http.MethodGet {
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}
		path := ctx.Param("filepath")

		fileName := strings.TrimSuffix(path, "/")
		if !strings.Contains(fileName, ".") {
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}
		ctx.Next()
	}
}
