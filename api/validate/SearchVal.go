package validate

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"io"
	"pledge-backend-test/api/common/stateCode"
	"pledge-backend-test/api/entity"
)

type SearchVal struct {
}

func NewSearchVal() *SearchVal {
	return &SearchVal{}
}

func (s *SearchVal) Search(ctx *gin.Context, search *entity.SearchEntity) int {
	err := ctx.ShouldBind(search)
	if err == io.EOF {
		return stateCode.ParamterEmptyErr
	} else if err != nil {
		errors := err.(validator.ValidationErrors)
		for _, e := range errors {
			if e.Field() == "chainID" && e.Tag() == "required" {
				return stateCode.ChainIdEmpty
			}
		}
		return stateCode.CommonErrServerErr
	}
	return stateCode.CommonSuccess
}
