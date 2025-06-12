package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
*跨域资源共享（CORS）中间件，用于处理浏览器跨域请求
 */
func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method               //请求方法（GET/POST等）
		origin := ctx.Request.Header.Get("Origin") //请求来源域名(如 https://example.com）
		// 设置 CORS 响应头
		if origin != "" {
			ctx.Header("Access-Control-Allow-Origin", "*")                                                                                                                          // 允许所有域名访问
			ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")                                                                                   // 允许的HTTP方法
			ctx.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, authCode, token, Content-Type, Accept, Authorization")                                            // 允许的请求头
			ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type") // 允许客户端访问的响应头
			ctx.Header("Access-Control-Allow-Credentials", "false")                                                                                                                 // 不允许携带凭据（如Cookies）
			ctx.Set("content-type", "application/json")                                                                                                                             // 强制响应类型为JSON
		}
		//处理预检请求（OPTIONS）
		if method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent) // 返回204状态码
		}
		//放行请求
		ctx.Next()
	}
}
