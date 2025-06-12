package validate

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"io"
	"pledge-backend-test/api/common/stateCode"
	"pledge-backend-test/api/entity"
)

type PoolBaseInfoValidate struct {
}

func NewPoolBaseInfoValidate() *PoolBaseInfoValidate {
	return &PoolBaseInfoValidate{}
}

func (p *PoolBaseInfoValidate) PoolBaseInfo(ctx *gin.Context, poolBaseInfo *entity.PoolBaseInfoEntity) int {
	//参数绑定
	err := ctx.ShouldBind(poolBaseInfo)
	if err == io.EOF {
		return stateCode.ParamterEmptyErr
	} else if err != nil {
		errors := err.(validator.ValidationErrors)
		for _, e := range errors {
			if e.Field() == "ChainId" && e.Tag() == "required" {
				return stateCode.ChainIdEmpty
			}
		}
		return stateCode.CommonErrServerErr
	}
	if poolBaseInfo.ChainId != 56 && poolBaseInfo.ChainId != 97 {
		return stateCode.ChainIdErr
	}
	return stateCode.CommonSuccess
}
