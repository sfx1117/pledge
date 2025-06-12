package service

import (
	"pledge-backend-test/api/common/stateCode"
	"pledge-backend-test/api/dao"
	"pledge-backend-test/api/models/response"
	"pledge-backend-test/log"
)

type PoolDataInfoService struct {
}

func NewPoolDataInfoService() *PoolDataInfoService {
	return &PoolDataInfoService{}
}

func (p *PoolDataInfoService) SelectPoolDataInfo(chainId int, data *[]response.PoolDataInfoRes) int {
	err := dao.NewPoolDataInfoDao().SelectPoolDataInfo(chainId, data)
	if err != nil {
		log.Logger.Error(err.Error())
		return stateCode.CommonErrServerErr
	}
	return stateCode.CommonSuccess
}
