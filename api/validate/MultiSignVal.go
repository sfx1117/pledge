package validate

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"io"
	"pledge-backend-test/api/common/stateCode"
	"pledge-backend-test/api/entity"
)

// 多签请求 参数校验
type MultiSignVal struct {
}

func NewMultiSignVal() *MultiSignVal {
	return &MultiSignVal{}
}

func (v *MultiSignVal) MultiSignEntityVal(ctx *gin.Context, req *entity.MultiSignEntity) int {
	//将 HTTP 请求的 JSON/表单数据绑定到 req 结构体
	err := ctx.ShouldBind(req)
	if req.ChainId != 97 && req.ChainId != 56 {
		return stateCode.ChainIdErr
	}
	//io.EOF：请求体为空
	if err == io.EOF {
		return stateCode.ParamterEmptyErr
	} else if err != nil { //参数验证失败
		errors := err.(validator.ValidationErrors)
		for _, e := range errors {
			if e.Field() == "SpName" && e.Tag() == "required" {
				return stateCode.PNameEmpty
			}
		}
		return stateCode.CommonErrServerErr
	}
	return stateCode.CommonSuccess
}

func (v *MultiSignVal) GetMultiSign(ctx *gin.Context, req *entity.GetMultiSign) int {
	//将 HTTP 请求的 JSON/表单数据绑定到 req 结构体
	err := ctx.ShouldBind(req)
	if req.ChainId != 97 && req.ChainId != 56 {
		return stateCode.ChainIdErr
	}
	//io.EOF：请求体为空
	if err == io.EOF {
		return stateCode.ParamterEmptyErr
	} else if err != nil { //参数验证失败
		errors := err.(validator.ValidationErrors)
		for _, e := range errors {
			if e.Field() == "SpName" && e.Tag() == "required" {
				return stateCode.PNameEmpty
			}
		}
		return stateCode.CommonErrServerErr
	}
	return stateCode.CommonSuccess
}
