package service

import (
	"fmt"
	"pledge-backend-test/api/common/stateCode"
	"pledge-backend-test/api/dao"
	"pledge-backend-test/api/entity"
	"pledge-backend-test/log"
)

type SearchService struct {
}

func NewSearchService() *SearchService {
	return &SearchService{}
}

func (s *SearchService) Search(req *entity.SearchEntity) (int, int64, []dao.PoolDao) {
	whereCondition := fmt.Sprintf(`chain_id='%v'`, req.ChainId)
	if req.LendTokenSymobl != "" {
		whereCondition += fmt.Sprintf(`and lend_token_symbol='%v'`, req.LendTokenSymobl)
	}
	if req.State != "" {
		whereCondition += fmt.Sprintf(`and state='%v'`, req.State)
	}
	err, total, poolList := dao.NewPoolDao().Pagination(req, whereCondition)
	if err != nil {
		log.Logger.Error(err.Error())
		return stateCode.CommonErrServerErr, 0, nil
	}
	return stateCode.CommonSuccess, total, poolList
}
