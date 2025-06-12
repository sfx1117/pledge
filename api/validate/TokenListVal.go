package validate

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"io"
	"pledge-backend-test/api/common/stateCode"
	"pledge-backend-test/api/entity"
)

type TokenListVal struct {
}

func NewTokenListVal() *TokenListVal {
	return &TokenListVal{}
}

func (t *TokenListVal) TokenListVal(ctx *gin.Context, tokenList *entity.TokenListEntity) int {
	err := ctx.ShouldBind(tokenList)
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
	if tokenList.ChainId != 97 && tokenList.ChainId != 56 {
		return stateCode.ChainIdErr
	}
	return stateCode.CommonSuccess
}
