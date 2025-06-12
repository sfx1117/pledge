package dao

import (
	"pledge-backend-test/db"
)

type TokenInfo struct {
	Id      int32  `json:"-" gorm:"column:id;primaryKey"`
	Symbol  string `json:"symbol" gorm:"column:symbol"`
	Token   string `json:"token" gorm:"column:token"`
	ChainId int    `json:"chain_id" gorm:"column:chain_id"`
}

type TokenList struct {
	Id           int    `json:"id" gorm:"cloumn:id;primaryKey"`
	Symbol       string `json:"symbol" gorm:"cloumn:symbol"`
	Logo         string `json:"logo" gorm:"cloumn:logo"`
	Price        string `json:"price" gorm:"cloumn:price"`
	Token        string `json:"symbol" gorm:"cloumn:token"`
	ChainId      int    `json:"chainId" gorm:"cloumn:chain_id"`
	AbiFileExist int    `json:"abiFileExist" gorm:"cloumn:abi_file_exist"`
	Decimals     int    `json:"decimals" gorm:"cloumn:decimals"`
}

func NewTokenListDao() *TokenList {
	return &TokenList{}
}

func (t *TokenList) GetTokenList(chainId int) (error, []TokenList) {
	var tokenList = make([]TokenList, 0)
	err := db.Mysql.Table("token_info").Where("chain_id=?", chainId).Find(&tokenList).Debug().Error
	if err != nil {
		return err, nil
	}
	return nil, tokenList
}

func (t *TokenList) DebtTokenList(chainId int) (error, []TokenInfo) {
	var tokenInfo = make([]TokenInfo, 0)
	err := db.Mysql.Table("token_info").Where("chain_id=?", chainId).Find(&tokenInfo).Debug().Error
	if err != nil {
		return err, nil
	}
	return nil, tokenInfo
}
