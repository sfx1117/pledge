package service

import (
	"pledge-backend-test/api/common/stateCode"
	"pledge-backend-test/api/dao"
	"pledge-backend-test/api/models/response"
	"pledge-backend-test/log"
)

type PoolBaseInfoService struct {
}

func NewPoolBaseInfoService() *PoolBaseInfoService {
	return &PoolBaseInfoService{}
}

func (p *PoolBaseInfoService) SelectPoolBaseInfo(chainId int, data *[]response.PoolBaseInfoRes) int {
	err := dao.NewPoolBaseInfoDao().SelectPoolBaseInfo(chainId, data)
	if err != nil {
		log.Logger.Error(err.Error())
		return stateCode.CommonErrServerErr
	}
	return stateCode.CommonSuccess
}
