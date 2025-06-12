package model

import (
	"errors"
	"gorm.io/gorm"
	"pledge-backend-test/db"
	"pledge-backend-test/log"
	"pledge-backend-test/utils"
)

type PoolBase struct {
	Id                     int    `json:"-" gorm:"column:id;primaryKey"`
	PoolId                 int    `json:"pool_id" gorm:"column:pool_id"`
	ChainId                string `json:"chain_id" gorm:"column:chain_id"`
	SettleTime             string `json:"settle_time" gorm:"column:settle_time"`
	EndTime                string `json:"end_time" gorm:"column:end_time"`
	InterestRate           string `json:"interest_rate" gorm:"column:interest_rate"`
	MaxSupply              string `json:"max_supply" gorm:"max_supply:"`
	LendSupply             string `json:"lend_supply" gorm:"column:lend_supply"`
	BorrowSupply           string `json:"borrow_supply" gorm:"column:borrow_supply"`
	MartgageRate           string `json:"martgage_rate" gorm:"column:martgage_rate"`
	LendToken              string `json:"lend_token" gorm:"column:lend_token"`
	LendTokenInfo          string `json:"lend_token_info" gorm:"column:lend_token_info"`
	BorrowToken            string `json:"borrow_token" gorm:"column:borrow_token"`
	BorrowTokenInfo        string `json:"borrow_token_info" gorm:"column:borrow_token_info"`
	State                  string `json:"state" gorm:"column:state"`
	SpCoin                 string `json:"sp_coin" gorm:"column:sp_coin"`
	JpCoin                 string `json:"jp_coin" gorm:"column:jp_coin"`
	LendTokenSymbol        string `json:"lend_token_symbol" gorm:"column:lend_token_symbol"`
	BorrowTokenSymbol      string `json:"borrow_token_symbol" gorm:"column:borrow_token_symbol"`
	AutoLiquidateThreshold string `json:"auto_liquidate_threshold" gorm:"column:auto_liquidate_threshold"`
	CreatedAt              string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt              string `json:"updated_at" gorm:"column:updated_at"`
}

type BorrowToken struct {
	BorrowFee  string `json:"borrowFee"`
	TokenLogo  string `json:"tokenLogo"`
	TokenName  string `json:"tokenName"`
	TokenPrice string `json:"tokenPrice"`
}

type LendToken struct {
	LendFee    string `json:"lendFee"`
	TokenLogo  string `json:"tokenLogo"`
	TokenName  string `json:"tokenName"`
	TokenPrice string `json:"tokenPrice"`
}

func NewPoolBase() *PoolBase {
	return &PoolBase{}
}

func (p *PoolBase) SavePoolBase(chainId string, poolId string, poolBase *PoolBase) error {
	//当前时间
	nowDateTime := utils.GetCurDateTimeFormat()
	//saveTokenInfo
	err, symbols := p.SaveTokenInfo(poolBase)
	if err != nil {
		log.Logger.Error(err.Error())
		return err
	}
	poolBase.BorrowTokenSymbol = symbols[0]
	poolBase.LendTokenSymbol = symbols[1]
	//数据库查询记录
	err = db.Mysql.Table("poolbases").Where("chain_id=? and pool_id=?", chainId, poolId).First(&p).Debug().Error
	if err != nil {
		//若无，则新增
		if errors.Is(err, gorm.ErrRecordNotFound) {
			poolBase.UpdatedAt = nowDateTime
			poolBase.CreatedAt = nowDateTime
			err = db.Mysql.Table("poolbases").Create(poolBase).Debug().Error
			if err != nil {
				log.Logger.Error(err.Error())
				return err
			}
		} else {
			return errors.New("record select err " + err.Error())
		}
	}
	//若有，则更新
	poolBase.UpdatedAt = nowDateTime
	err = db.Mysql.Table("poolbases").Where("chain_id=? and pool_id=?", chainId, poolId).Updates(poolBase).Debug().Error
	if err != nil {
		log.Logger.Error(err.Error())
		return err
	}
	return nil
}

func (p *PoolBase) SaveTokenInfo(poolBase *PoolBase) (error, []string) {
	tokenInfo := TokenInfo{}
	symbols := []string{"", ""}

	//当前时间
	nowDataTime := utils.GetCurDateTimeFormat()

	//borrowTokeninfo
	err := db.Mysql.Table("token_info").Where("chain_id=? and token=?", poolBase.ChainId, poolBase.BorrowToken).First(&tokenInfo).Debug().Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tokenInfo.Token = poolBase.BorrowToken
			err = db.Mysql.Table("token_info").Create(&TokenInfo{
				Token:     poolBase.BorrowToken,
				ChainId:   poolBase.ChainId,
				UpdatedAt: nowDataTime,
				CreatedAt: nowDataTime,
			}).Debug().Error
			if err != nil {
				log.Logger.Error(err.Error())
				return err, symbols
			}
		} else {
			return errors.New("token_info record select err " + err.Error()), symbols
		}
	}
	symbols[0] = tokenInfo.Symbol

	//lendTokeninfo
	tokenInfo = TokenInfo{}
	err = db.Mysql.Table("token_info").Where("chain_id=? and token=?", poolBase.ChainId, poolBase.LendToken).First(&tokenInfo).Debug().Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tokenInfo.Token = poolBase.LendToken
			err = db.Mysql.Table("token_info").Create(&TokenInfo{
				Token:     poolBase.LendToken,
				ChainId:   poolBase.ChainId,
				UpdatedAt: nowDataTime,
				CreatedAt: nowDataTime,
			}).Debug().Error
			if err != nil {
				log.Logger.Error(err.Error())
				return err, symbols
			}
		} else {
			return errors.New("token_info record select err " + err.Error()), symbols
		}
	}
	symbols[1] = tokenInfo.Symbol
	return nil, symbols
}
