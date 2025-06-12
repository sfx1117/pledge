package dao

import (
	"pledge-backend-test/api/models/response"
	"pledge-backend-test/db"
)

type PoolDataInfoDao struct {
	PoolId                 int    `json:"poolId" gorm:"cloumn:pool_id"`
	SettleAmountLend       string `json:"settleAmountLend" gorm:"cloumn:settle_amount_lend"`
	SettleAmountBorrow     string `json:"settleAmountBorrow" gorm:"cloumn:settle_amount_borrow"`
	FinishAmountLend       string `json:"finishAmountLend" gorm:"cloumn:finish_amount_lend"`
	FinishAmountBorrow     string `json:"finishAmountBorrow" gorm:"cloumn:finish_amount_borrow"`
	LiquidationAmounLend   string `json:"liquidationAmounLend" gorm:"cloumn:liquidation_amoun_lend"`
	LiquidationAmounBorrow string `json:"liquidationAmounBorrow" gorm:"cloumn:liquidation_amoun_borrow"`
	ChainId                string `json:"chainId" gorm:"cloumn:chain_id"`
}

func NewPoolDataInfoDao() *PoolDataInfoDao {
	return &PoolDataInfoDao{}
}

func (p *PoolDataInfoDao) SelectPoolDataInfo(chainId int, data *[]response.PoolDataInfoRes) error {
	var poolDataInfos []PoolDataInfoDao
	err := db.Mysql.Table("pooldata").Where("chain_id=?", chainId).Order("pool_id asc").Find(&poolDataInfos).Debug().Error
	if err != nil {
		return err
	}
	for _, poolData := range poolDataInfos {
		*data = append(*data, response.PoolDataInfoRes{
			Index:                  poolData.PoolId - 1,
			PoolId:                 poolData.PoolId,
			SettleAmountLend:       poolData.SettleAmountLend,
			SettleAmountBorrow:     poolData.SettleAmountBorrow,
			FinishAmountLend:       poolData.FinishAmountLend,
			FinishAmountBorrow:     poolData.FinishAmountBorrow,
			LiquidationAmounLend:   poolData.LiquidationAmounLend,
			LiquidationAmounBorrow: poolData.LiquidationAmounBorrow,
			ChainId:                poolData.ChainId,
		})
	}
	return nil
}
