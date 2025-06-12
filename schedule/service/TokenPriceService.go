package service

import (
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"pledge-backend-test/config"
	"pledge-backend-test/contract/bindings"
	"pledge-backend-test/db"
	"pledge-backend-test/log"
	"pledge-backend-test/schedule/model"
	"pledge-backend-test/utils"
	"strings"
)

type TokenPriceService struct {
}

func NewTokenPriceService() *TokenPriceService {
	return &TokenPriceService{}
}

// 更新合约价格
func (s *TokenPriceService) UpdateContractPrice() {
	tokens := []model.TokenInfo{}
	db.Mysql.Table("token_info").Find(&tokens)
	for _, t := range tokens {
		var err error
		var price int64 = 0

		if t.Token == "" {
			log.Logger.Sugar().Error("UpdateContractPrice token empty ", t.Symbol, t.ChainId)
			continue
		} else {
			if t.ChainId == "97" {
				err, price = s.GetTestNetTokenPrice(t.Token)
			} else if t.ChainId == "56" {
				if strings.ToUpper(t.Token) == config.Config.MainNet.PlgrAddress {
					priceStr, _ := db.RedisGetString("pledge_price")
					priceDec, _ := decimal.NewFromString(priceStr)
					e8 := decimal.NewFromInt(100000000)
					priceDec = priceDec.Mul(e8)
					price = priceDec.IntPart()
				} else {
					err, price = s.GetMainNetTokenPrice(t.Token)
				}
			}
			if err != nil {
				log.Logger.Sugar().Error("UpdateContractPrice err ", t.Symbol, t.ChainId, err)
				continue
			}
		}

		hasNewData, err := s.CheckPriceData(t.Token, t.ChainId, utils.Int64ToString(price))
		if err != nil {
			log.Logger.Sugar().Error("UpdateContractPrice CheckPriceData err ", err)
			continue
		}
		if hasNewData {
			err = s.UpdateTokenInfoPrice(t.Token, t.ChainId, utils.Int64ToString(price))
			if err != nil {
				log.Logger.Sugar().Error("UpdateContractPrice SavePriceData err ", err)
				continue
			}
		}
	}

}

// 获取测试网token价格
func (s *TokenPriceService) GetTestNetTokenPrice(token string) (error, int64) {
	client, err := ethclient.Dial(config.Config.TestNet.NetUrl)
	if err != nil {
		log.Logger.Sugar().Error(err.Error())
		return err, 0
	}
	testnetToken, err := bindings.NewBscPledgeOracleTestnetToken(common.HexToAddress(config.Config.TestNet.BscPledgeOracleToken), client)
	if err != nil {
		log.Logger.Sugar().Error(err.Error())
		return err, 0
	}
	price, err := testnetToken.GetPrice(nil, common.HexToAddress(token))
	if err != nil {
		log.Logger.Sugar().Error(err.Error())
		return err, 0
	}
	return nil, price.Int64()
}

// 获取主网token价格
func (s *TokenPriceService) GetMainNetTokenPrice(token string) (error, int64) {
	client, err := ethclient.Dial(config.Config.MainNet.NetUrl)
	if err != nil {
		log.Logger.Sugar().Error(err.Error())
		return err, 0
	}
	mainnetToken, err := bindings.NewBscPledgeOracleMainnetToken(common.HexToAddress(config.Config.MainNet.BscPledgeOracleToken), client)
	if err != nil {
		log.Logger.Sugar().Error(err.Error())
		return err, 0
	}
	price, err := mainnetToken.GetPrice(nil, common.HexToAddress(token))
	if err != nil {
		log.Logger.Sugar().Error(err.Error())
		return err, 0
	}
	return nil, price.Int64()
}

func (s *TokenPriceService) CheckPriceData(token string, chainId string, price string) (bool, error) {
	redisKey := "tokenInfo:" + chainId + ":" + token
	tokenInfoBytes, err := db.RedisGet(redisKey)
	if len(tokenInfoBytes) <= 0 {
		err = s.CheckTokenInfo(token, chainId)
		if err != nil {
			log.Logger.Error(err.Error())
		}
		err = db.RedisSet(redisKey, model.RedisTokenInfo{
			Token:   token,
			ChainId: chainId,
			Price:   price,
		}, 0)
		if err != nil {
			log.Logger.Error(err.Error())
			return false, err
		}
	} else {
		redisTokenInfo := model.RedisTokenInfo{}
		err = json.Unmarshal(tokenInfoBytes, redisTokenInfo)
		if err != nil {
			log.Logger.Sugar().Error(err.Error())
			return false, err
		}
		if redisTokenInfo.Price == price {
			return false, nil
		}
		redisTokenInfo.Price = price
		err = db.RedisSet(redisKey, redisTokenInfo, 0)
		if err != nil {
			log.Logger.Sugar().Error(err.Error())
			return false, err
		}
	}
	return true, nil
}

func (s *TokenPriceService) UpdateTokenInfoPrice(token string, chainId string, price string) error {
	nowDataTime := utils.GetCurDateTimeFormat()
	err := db.Mysql.Table("token_info").Where("token=? and chain_id=?", token, chainId).Updates(map[string]interface{}{
		"price":      price,
		"updated_at": nowDataTime,
	}).Debug().Error
	if err != nil {
		log.Logger.Sugar().Error("UpdateContractPrice SavePriceData err ", err)
		return err
	}
	return nil
}

func (s *TokenPriceService) CheckTokenInfo(token string, chainId string) error {
	tokenInfo := model.TokenInfo{}
	err := db.Mysql.Table("token_info").Where("token=? and chain_id=?", token, chainId).First(&tokenInfo).Debug().Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tokenInfo = model.TokenInfo{}
			nowDataTime := utils.GetCurDateTimeFormat()
			tokenInfo.Token = token
			tokenInfo.ChainId = chainId
			tokenInfo.UpdatedAt = nowDataTime
			tokenInfo.CreatedAt = nowDataTime
			err = db.Mysql.Table("token_info").Create(tokenInfo).Debug().Error
			if err != nil {
				log.Logger.Sugar().Error(err.Error())
				return err
			}
		} else {
			return err
		}
	}
	return nil
}
