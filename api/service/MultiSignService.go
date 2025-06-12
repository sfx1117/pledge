package service

import (
	"encoding/json"
	"pledge-backend-test/api/common/stateCode"
	"pledge-backend-test/api/dao"
	"pledge-backend-test/api/entity"
	"pledge-backend-test/api/models/response"
)

type MultiSignService struct {
}

func NewMultiSignService() *MultiSignService {
	return &MultiSignService{}
}

// 设置多签数据
func (ms *MultiSignService) SetMultiSign(multiSign *entity.MultiSignEntity) (int, error) {
	err := dao.NewMultiSignDao().Set(multiSign)
	if err != nil {
		return stateCode.CommonErrServerErr, err
	}
	return stateCode.CommonSuccess, nil
}

func (ms *MultiSignService) GetMultiSign(data *response.MultiSignRes, chainId int) (int, error) {
	multiSignDao := dao.NewMultiSignDao()
	err := multiSignDao.Get(chainId)
	if err != nil {
		return stateCode.CommonSuccess, err
	}
	//将多签账号转换为[]string
	var multiSignAccount []string
	_ = json.Unmarshal([]byte(multiSignDao.MultiSignAccount), &multiSignAccount)
	//数据组装
	data.MultiSignAccount = multiSignAccount
	data.SpName = multiSignDao.SpName
	data.SpToken = multiSignDao.SpToken
	data.SpAddress = multiSignDao.SpAddress
	data.SpHash = multiSignDao.SpHash
	data.JpName = multiSignDao.JpName
	data.JpToken = multiSignDao.JpToken
	data.JpAddress = multiSignDao.JpAddress
	data.JpHash = multiSignDao.JpHash

	return stateCode.CommonSuccess, nil
}
