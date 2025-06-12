package validate

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"io"
	"pledge-backend-test/api/common/stateCode"
	"pledge-backend-test/api/entity"
)

type PoolDataInfoVal struct {
}

func NewPoolDataInfoVal() *PoolDataInfoVal {
	return &PoolDataInfoVal{}
}

func (p *PoolDataInfoVal) PoolDataInfo(ctx *gin.Context, poolData *entity.PoolDataInfoEntity) int {
	err := ctx.ShouldBind(poolData)
	if err == io.EOF {
		return stateCode.ParamterEmptyErr
	} else if err != nil {
		errors := err.(validator.ValidationErrors)
		for _, e := range errors {
			if e.Field() == "chainId" && e.Tag() == "required" {
				return stateCode.ChainIdEmpty
			}
		}
		return stateCode.CommonErrServerErr
	}
	if poolData.ChainId != 56 && poolData.ChainId != 97 {
		return stateCode.ChainIdErr
	}
	return stateCode.CommonSuccess
}
