package controller

import (
	"github.com/gin-gonic/gin"
	"pledge-backend-test/api/common/stateCode"
	"pledge-backend-test/api/entity"
	"pledge-backend-test/api/models/response"
	"pledge-backend-test/api/service"
	"pledge-backend-test/api/validate"
	"pledge-backend-test/db"
)

type UserController struct {
}

/*
*
登录
*/
func (u *UserController) Login(ctx *gin.Context) {
	res := response.Gin{Res: ctx}
	req := entity.UserEntity{}
	data := response.UserRes{}

	errcode := validate.NewUserVal().Login(ctx, &req)
	if errcode != stateCode.CommonSuccess {
		res.Response(ctx, errcode, nil)
		return
	}
	errcode = service.NewUserService().Login(&req, &data)
	if errcode != stateCode.CommonSuccess {
		res.Response(ctx, errcode, nil)
		return
	}
	res.Response(ctx, stateCode.CommonSuccess, data)
	return
}

/*
*
登出
*/
func (u *UserController) Logout(ctx *gin.Context) {
	res := response.Gin{Res: ctx}
	userName, _ := ctx.Get("userName")
	//删除缓存
	_, _ = db.RedisDelete(userName.(string))
	res.Response(ctx, stateCode.CommonSuccess, nil)
	return
}
