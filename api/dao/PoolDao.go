package dao

import (
	"encoding/json"
	"pledge-backend-test/api/entity"
	"pledge-backend-test/api/models/response"
	"pledge-backend-test/db"
)

type SearchRes struct {
	Count int64     `json:"count"`
	Rows  []PoolDao `json:"rows"`
}

type PoolDao struct {
	PoolID                 int             `json:"pool_id"`
	SettleTime             string          `json:"settleTime"`
	EndTime                string          `json:"endTime"`
	InterestRate           string          `json:"interestRate"`
	MaxSupply              string          `json:"maxSupply"`
	LendSupply             string          `json:"lendSupply"`
	BorrowSupply           string          `json:"borrowSupply"`
	MartgageRate           string          `json:"martgageRate"`
	LendToken              string          `json:"lendToken"`
	LendTokenSymbol        string          `json:"lend_token_symbol"`
	BorrowToken            string          `json:"borrowToken"`
	BorrowTokenSymbol      string          `json:"borrow_token_symbol"`
	State                  string          `json:"state"`
	SpCoin                 string          `json:"spCoin"`
	JpCoin                 string          `json:"jpCoin"`
	AutoLiquidateThreshold string          `json:"autoLiquidateThreshold"`
	Pooldata               PoolDataInfoDao `json:"pooldata"`
}

func NewPoolDao() *PoolDao {
	return &PoolDao{}
}

func (p *PoolDao) Pagination(req *entity.SearchEntity, whereCondition string) (error, int64, []PoolDao) {
	var total int64
	var pools []PoolDao
	var poolBases []PoolBase
	var poolData PoolDataInfoDao

	db.Mysql.Table("poolbases").Where(whereCondition).Count(&total)
	//分页查询
	err := db.Mysql.Table("poolbases").Where(whereCondition).Order("pool_id asc").Limit(req.PageSize).Offset((req.Page - 1) * req.PageSize).Find(&poolBases).Debug().Error
	if err != nil {
		return err, 0, nil
	}
	//poolData
	err = db.Mysql.Table("pooldata").Where("chain_id=?", req.ChainId).First(&poolData).Debug().Error
	if err != nil {
		return err, 0, nil
	}
	//数据组装
	for _, v := range poolBases {
		var borrowTokenInfo response.BorrowTokenInfo
		json.Unmarshal([]byte(v.BorrowTokenInfo), borrowTokenInfo)
		var lendTokenInfo response.LendTokenInfo
		json.Unmarshal([]byte(v.LendTokenInfo), lendTokenInfo)

		pools = append(pools, PoolDao{
			PoolID:                 v.PoolId,
			SettleTime:             v.SettleTime,
			EndTime:                v.EndTime,
			InterestRate:           v.InterestRate,
			MaxSupply:              v.MaxSupply,
			LendSupply:             v.LendSupply,
			BorrowSupply:           v.BorrowSupply,
			MartgageRate:           v.MartgageRate,
			LendToken:              lendTokenInfo.TokenName,
			BorrowToken:            borrowTokenInfo.TokenName,
			State:                  v.State,
			SpCoin:                 v.SpCoin,
			JpCoin:                 v.JpCoin,
			AutoLiquidateThreshold: v.AutoLiquidateThreshold,
			Pooldata:               poolData})

	}
	return nil, total, pools
}
