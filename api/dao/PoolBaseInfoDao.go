package dao

import (
	"encoding/json"
	"pledge-backend-test/api/models/response"
	"pledge-backend-test/db"
)

// 表结构
type PoolBase struct {
	PoolId                 int    `json:"poolId" gorm:"cloumn:pool_id"`
	SettleTime             string `json:"settleTime" gorm:"cloumn:settle_time"`
	EndTime                string `json:"endTime" gorm:"cloumn:end_time"`
	InterestRate           string `json:"interestRate" gorm:"cloumn:interest_rate"`
	MaxSupply              string `json:"maxSupply" gorm:"cloumn:max_supply"`
	LendSupply             string `json:"lendSupply" gorm:"cloumn:lend_supply"`
	BorrowSupply           string `json:"borrowSupply" gorm:"cloumn:borrow_supply"`
	MartgageRate           string `json:"martgageRate" gorm:"cloumn:martgage_rate"`
	LendToken              string `json:"lendToken" gorm:"cloumn:lend_token"`
	BorrowToken            string `json:"borrowToken" gorm:"cloumn:borrow_token"`
	State                  string `json:"state" gorm:"cloumn:state"`
	JpCoin                 string `json:"jpCoin" gorm:"cloumn:jp_coin"`
	SpCoin                 string `json:"spCoin" gorm:"cloumn:sp_coin"`
	AutoLiquidateThreshold string `json:"autoLiquidateThreshold" gorm:"cloumn:auto_liquidate_threshold"`
	BorrowTokenInfo        string `json:"borrowTokenInfo" gorm:"cloumn:borrow_token_info"`
	LendTokenInfo          string `json:"lendTokenInfo" gorm:"cloumn:lend_token_info"`
	ChainId                string `json:"chainId" gorm:"cloumn:chain_id"`
	LendTokenSymbol        string `json:"lendTokenSymbol" gorm:"cloumn:lend_token_symbol"`
	BorrowTokenSymbol      string `json:"borrowTokenSymbol" gorm:"cloumn:borrow_token_symbol"`
}

func NewPoolBaseInfoDao() *PoolBase {
	return &PoolBase{}
}

func (p *PoolBase) SelectPoolBaseInfo(chainId int, data *[]response.PoolBaseInfoRes) error {
	var poolBases []PoolBase
	err := db.Mysql.Table("poolbases").Where("chain_id=?", chainId).Order("pool_id asc").Find(&poolBases).Debug().Error
	if err != nil {
		return err
	}
	//解析返回参数
	for _, poolBase := range poolBases {
		borrowTokenInfo := response.BorrowTokenInfo{}
		_ = json.Unmarshal([]byte(poolBase.BorrowTokenInfo), &borrowTokenInfo)
		lendTokenInfo := response.LendTokenInfo{}
		_ = json.Unmarshal([]byte(poolBase.LendTokenInfo), &lendTokenInfo)
		*data = append(*data, response.PoolBaseInfoRes{
			Index:                  poolBase.PoolId - 1,
			PoolId:                 poolBase.PoolId,
			SettleTime:             poolBase.SettleTime,
			InterestRate:           poolBase.InterestRate,
			MaxSupply:              poolBase.MaxSupply,
			LendSupply:             poolBase.LendSupply,
			BorrowSupply:           poolBase.BorrowSupply,
			MartgageRate:           poolBase.MartgageRate,
			LendToken:              poolBase.LendToken,
			BorrowToken:            poolBase.BorrowToken,
			State:                  poolBase.State,
			JpCoin:                 poolBase.JpCoin,
			SpCoin:                 poolBase.SpCoin,
			AutoLiquidateThreshold: poolBase.AutoLiquidateThreshold,
			BorrowTokenInfo:        borrowTokenInfo,
			LendTokenInfo:          lendTokenInfo,
			ChainId:                poolBase.ChainId,
			LendTokenSymbol:        poolBase.LendTokenSymbol,
			BorrowTokenSymbol:      poolBase.BorrowTokenSymbol,
		})
	}
	return nil
}
