package middleware

import (
	"github.com/gin-gonic/gin"
	"pledge-backend-test/api/common/stateCode"
	"pledge-backend-test/api/models/response"
	"pledge-backend-test/config"
	"pledge-backend-test/db"
	"pledge-backend-test/utils"
)

func CheckToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res := response.Gin{Res: ctx}
		//从请求头中获取token
		token := ctx.Request.Header.Get("authCode")
		//用jwt将token解析
		userName, err := utils.ParseToken(token, config.Config.Jwt.SecretKey)
		if err != nil {
			res.Response(ctx, stateCode.TokenErr, nil)
			ctx.Abort()
			return
		}
		//j校验是否默认username
		if userName != config.Config.DefaultAdmin.UserName {
			res.Response(ctx, stateCode.TokenErr, nil)
			ctx.Abort()
			return
		}
		//校验是否登录成功
		resByteArr, err := db.RedisGet(userName)
		if string(resByteArr) != `"login_ok"` {
			res.Response(ctx, stateCode.TokenErr, nil)
			ctx.Abort()
			return
		}

		ctx.Set("userName", userName)
		ctx.Next()
	}
}
