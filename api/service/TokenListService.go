package service

import (
	"pledge-backend-test/api/common/stateCode"
	"pledge-backend-test/api/dao"
	"pledge-backend-test/log"
)

type TokenListService struct {
}

func NewTokenListService() *TokenListService {
	return &TokenListService{}
}

func (t *TokenListService) GetTokenList(chainId int) (int, []dao.TokenList) {
	err, tokenList := dao.NewTokenListDao().GetTokenList(chainId)
	if err != nil {
		log.Logger.Error(err.Error())
		return stateCode.CommonErrServerErr, nil
	}
	return stateCode.CommonSuccess, tokenList
}

func (t *TokenListService) DebtTokenList(chainId int) (int, []dao.TokenInfo) {
	err, tokenInfo := dao.NewTokenListDao().DebtTokenList(chainId)
	if err != nil {
		log.Logger.Error(err.Error())
		return stateCode.CommonErrServerErr, nil
	}
	return stateCode.CommonSuccess, tokenInfo
}
