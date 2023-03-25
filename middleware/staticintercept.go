package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func StaticInterceptor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 拦截get请求
		if ctx.Request.Method == http.MethodGet {
			path := ctx.Param("filepath")
			fileName := strings.TrimSuffix(path, "/")
			if !strings.Contains(fileName, ".") {
				ctx.AbortWithStatus(http.StatusForbidden)
				return
			}
		}
		ctx.Next()
	}
}
