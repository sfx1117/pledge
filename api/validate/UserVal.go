package validate

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"io"
	"pledge-backend-test/api/common/stateCode"
	"pledge-backend-test/api/entity"
)

type UserVal struct {
}

func NewUserVal() *UserVal {
	return &UserVal{}
}

func (u *UserVal) Login(ctx *gin.Context, req *entity.UserEntity) int {
	err := ctx.ShouldBind(req)
	if err == io.EOF {
		return stateCode.ParamterEmptyErr
	} else if err != nil {
		errors := err.(validator.ValidationErrors)
		for _, e := range errors {
			if e.Field() == "Name" && e.Tag() == "required" {
				return stateCode.PNameEmpty
			}
			if e.Field() == "Password" && e.Tag() == "required" {
				return stateCode.PNameEmpty
			}
		}
		return stateCode.CommonErrServerErr
	}
	return stateCode.CommonSuccess
}
