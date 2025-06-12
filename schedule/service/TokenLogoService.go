package service

import (
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"pledge-backend-test/config"
	"pledge-backend-test/db"
	"pledge-backend-test/log"
	"pledge-backend-test/schedule/model"
	"pledge-backend-test/utils"
	"regexp"
	"strings"
)

type TokenLogoService struct {
}

func NewTokenLogoService() *TokenLogoService {
	return &TokenLogoService{}
}

func (s *TokenLogoService) UpdateTokenLogo() {
	res, err := utils.HttpGet(config.Config.Token.LogoUrl, map[string]string{})
	if err != nil {
		log.Logger.Sugar().Error("UpdateTokenLogo HttpGet err", err)
	} else {
		tokenLogoRemote := model.TokenLogoRemote{}
		err = json.Unmarshal(res, &tokenLogoRemote)
		if err != nil {
			log.Logger.Sugar().Error("UpdateTokenLogo json err", err)
			return
		}

		for _, t := range tokenLogoRemote.Tokens {
			hasNewData, err := s.CheckLogoData(t.Address, utils.IntToString(t.ChainId), t.LogoURI, t.Symbol)
			if err != nil {
				log.Logger.Sugar().Error("UpdateTokenLogo CheckLogoData err", err)
				continue
			}
			if hasNewData {
				err = s.SaveLogoData(t.Address, utils.IntToString(t.ChainId), t.LogoURI, t.Symbol, t.Decimals)
				if err != nil {
					log.Logger.Sugar().Error("UpdateTokenLogo SaveLogoData err", err)
					continue
				}
			}
		}
	}

	for _, v := range LocalTokenLogo {
		for _, t := range v {
			if t["token"] == "" {
				continue
			}
			hasNewData, err := s.CheckLogoData(t["token"], t["chain_id"], t["logo"], t["symbol"])
			if err != nil {
				log.Logger.Sugar().Error("UpdateTokenLogo CheckLogoData err", err)
				continue
			}
			if hasNewData {
				err = s.SaveLogoData(t["token"], t["chain_id"], t["logo"], t["symbol"], utils.StringToInt(t["decimals"]))
				if err != nil {
					log.Logger.Sugar().Error("UpdateTokenLogo SaveLogoData err", err)
					continue
				}
			}
		}
	}

}

func (s *TokenLogoService) CheckLogoData(token, chainId, logoUrl, symbol string) (bool, error) {
	redisKey := "token_info:" + chainId + ":" + token
	tokenInfoBytes, _ := db.RedisGet(redisKey)
	if len(tokenInfoBytes) <= 0 {
		err := s.CheckTokenInfo(token, chainId)
		if err != nil {
			log.Logger.Sugar().Error(err.Error())
			return false, err
		}
		err = db.RedisSet(redisKey, model.RedisTokenInfo{
			Logo:    logoUrl,
			Token:   token,
			ChainId: chainId,
			Symbol:  symbol,
		}, 0)
		if err != nil {
			log.Logger.Sugar().Error(err.Error())
			return false, err
		}
	} else {
		redisTokenInfo := model.RedisTokenInfo{}
		err := json.Unmarshal(tokenInfoBytes, &redisTokenInfo)
		if err != nil {
			log.Logger.Sugar().Error(err.Error())
			return false, err
		}
		if redisTokenInfo.Logo == logoUrl {
			return false, err
		}
		redisTokenInfo.Logo = logoUrl
		redisTokenInfo.Symbol = symbol
		err = db.RedisSet(redisKey, redisTokenInfo, 0)
		if err != nil {
			log.Logger.Sugar().Error(err.Error())
			return false, err
		}
	}
	return true, nil
}
func (s *TokenLogoService) CheckTokenInfo(token string, chainId string) error {
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
				return err
			}
		} else {
			return err
		}
	}
	return nil
}
func (s *TokenLogoService) SaveLogoData(token, chainId, logoUrl, symbol string, decimals int) error {
	nowDataTime := utils.GetCurDateTimeFormat()
	err := db.Mysql.Table("token_info").Where("chain_id=? and token=?", chainId, token).Updates(map[string]interface{}{
		"logo":       logoUrl,
		"symbol":     symbol,
		"decimals":   decimals,
		"updated_at": nowDataTime,
	}).Debug().Error
	if err != nil {
		return err
	}
	return nil
}

func GetBaseUrl() string {
	domainName := config.Config.Env.DomainName
	domainNameSlice := strings.Split(domainName, "")
	pattern := "\\d+"
	isNumber, _ := regexp.MatchString(pattern, domainNameSlice[0])
	if isNumber {
		return config.Config.Env.Protocol + "://" + config.Config.Env.DomainName + ":" + config.Config.Env.Port + "/"
	}
	return config.Config.Env.Protocol + "://" + config.Config.Env.DomainName + "/"
}

var BaseUrl = GetBaseUrl()
var LocalTokenLogo = map[string]map[string]map[string]string{
	"BNB": {
		"test_net": {
			"chain_id": "97",
			"decimals": "18",
			"token":    "0x0000000000000000000000000000000000000000",
			"symbol":   "BNB",
			"logo":     BaseUrl + "static/img/BNB.png",
		},
		"main_net": {
			"chain_id": "56",
			"decimals": "18",
			"token":    "0x0000000000000000000000000000000000000000",
			"symbol":   "BNB",
			"logo":     BaseUrl + "static/img/BNB.png",
		},
	},
	"BTC": {
		"test_net": {
			"chain_id": "97",
			"decimals": "8",
			"token":    "0xB5514a4FA9dDBb48C3DE215Bc9e52d9fCe2D8658",
			"symbol":   "BTC",
			"logo":     BaseUrl + "static/img/BTC.png",
		},
		"main_net": {
			"chain_id": "56",
			"decimals": "8",
			"token":    "0x7130d2A12B9BCbFAe4f2634d864A1Ee1Ce3Ead9c",
			"symbol":   "BTC",
			"logo":     BaseUrl + "static/img/BTC.png",
		},
	},
	"BTCB": {
		"test_net": {
			"chain_id": "97",
			"decimals": "8",
			"token":    "0xB5514a4FA9dDBb48C3DE215Bc9e52d9fCe2D8658",
			"symbol":   "BTC",
			"logo":     BaseUrl + "static/img/BTC.png",
		},
		"main_net": {
			"chain_id": "56",
			"decimals": "8",
			"token":    "0x7130d2A12B9BCbFAe4f2634d864A1Ee1Ce3Ead9c",
			"symbol":   "BTC",
			"logo":     BaseUrl + "static/img/BTC.png",
		},
	},
	"BUSD": {
		"test_net": {
			"chain_id": "97",
			"decimals": "18",
			"token":    "0xE676Dcd74f44023b95E0E2C6436C97991A7497DA",
			"symbol":   "BUSD",
			"logo":     BaseUrl + "static/img/BUSD.png",
		},
		"main_net": {
			"chain_id": "56",
			"decimals": "18",
			"token":    "0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56",
			"symbol":   "BUSD",
			"logo":     BaseUrl + "static/img/BUSD.png",
		},
	},
	"DAI": {
		"test_net": {
			"chain_id": "97",
			"decimals": "18",
			"token":    "0x490BC3FCc845d37C1686044Cd2d6589585DE9B8B",
			"symbol":   "DAI",
			"logo":     BaseUrl + "static/img/DAI.png",
		},
		"main_net": {
			"chain_id": "56",
			"decimals": "18",
			"token":    "0x1AF3F329e8BE154074D8769D1FFa4eE058B1DBc3",
			"symbol":   "DAI",
			"logo":     BaseUrl + "static/img/DAI.png",
		},
	},
	"ETH": {
		"test_net": {
			"chain_id": "97",
			"decimals": "18",
			"token":    "",
			"symbol":   "ETH",
			"logo":     BaseUrl + "static/img/ETH.png",
		},
		"main_net": {
			"chain_id": "56",
			"decimals": "18",
			"token":    "0x2170ed0880ac9a755fd29b2688956bd959f933f8",
			"symbol":   "ETH",
			"logo":     BaseUrl + "static/img/ETH.png",
		},
	},
	"USDT": {
		"test_net": {
			"chain_id": "97",
			"decimals": "18",
			"token":    "",
			"symbol":   "USDT",
			"logo":     BaseUrl + "static/img/USDT.png",
		},
		"main_net": {
			"chain_id": "56",
			"decimals": "18",
			"token":    "0x55d398326f99059ff775485246999027b3197955",
			"symbol":   "USDT",
			"logo":     BaseUrl + "static/img/USDT.png",
		},
	},
	"CAKE": {
		"test_net": {
			"chain_id": "97",
			"decimals": "18",
			"token":    "0xEAEd08168a2D34Ae2B9ea1c1f920E0BC00F9fA67",
			"symbol":   "CAKE",
			"logo":     BaseUrl + "static/img/CAKE.png",
		},
		"main_net": {
			"chain_id": "56",
			"decimals": "18",
			"token":    "0x0e09fabb73bd3ade0a17ecc321fd13a19e81ce82",
			"symbol":   "CAKE",
			"logo":     BaseUrl + "static/img/CAKE.png",
		},
	},
	"PLGR": {
		"test_net": {
			"chain_id": "97",
			"decimals": "18",
			"token":    "",
			"symbol":   "PLGR",
			"logo":     BaseUrl + "static/img/PLGR.png",
		},
		"main_net": {
			"chain_id": "56",
			"decimals": "18",
			"token":    "0x6Aa91CbfE045f9D154050226fCc830ddbA886CED",
			"symbol":   "PLGR",
			"logo":     BaseUrl + "static/img/PLGR.png",
		},
	},
}
