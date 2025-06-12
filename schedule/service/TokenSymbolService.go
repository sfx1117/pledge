package service

import (
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
	"os"
	"pledge-backend-test/config"
	abifile "pledge-backend-test/contract/abi"
	"pledge-backend-test/db"
	"pledge-backend-test/log"
	"pledge-backend-test/schedule/model"
	"pledge-backend-test/utils"
	"strings"
)

type TokenSymbolService struct {
}

func NewTokenSymbolService() *TokenSymbolService {
	return &TokenSymbolService{}
}

func (s *TokenSymbolService) UpdateTokenSymbol() {
	tokens := []model.TokenInfo{}
	db.Mysql.Table("token_info").Find(&tokens)

	for _, t := range tokens {
		var err error
		var symbol string
		if t.Token == "" {
			log.Logger.Sugar().Error("UpdateContractSymbol token empty", t.Symbol, t.ChainId)
			continue
		}
		if t.ChainId == "97" {
			err, symbol = s.GetTestNetTokenSymbol(t.Token)
		} else if t.ChainId == "56" {
			if t.AbiFileExist == 0 {
				err = s.GetRemoteAbiFileByToken(t.Token, t.ChainId)
				if err != nil {
					log.Logger.Sugar().Error("UpdateContractSymbol GetRemoteAbiFileByToken err", t.Symbol, t.Symbol)
					continue
				}
			}
			err, symbol = s.GetMainNetTokenSymbol(t.Token)
		} else {
			log.Logger.Sugar().Error("UpdateContractSymbol chain_id err ", t.Symbol, t.ChainId)
			continue
		}
		if err != nil {
			log.Logger.Sugar().Error("UpdateContractSymbol err ", t.Symbol, t.ChainId, err)
			continue
		}

		hasNewData, err := s.CheckSymbolData(t.Token, t.ChainId, symbol)
		if err != nil {
			log.Logger.Sugar().Error("UpdateContractSymbol CheckSymbolData err ", err)
			continue
		}
		if hasNewData {
			err = s.SaveSymbolData(t.Token, t.ChainId, t.Symbol)
			if err != nil {
				log.Logger.Sugar().Error("UpdateContractSymbol CheckSymbolData err ", err)
				continue
			}
		}
	}

}

// 测试网
func (s *TokenSymbolService) GetTestNetTokenSymbol(token string) (error, string) {
	client, err := ethclient.Dial(config.Config.TestNet.NetUrl)
	if err != nil {
		log.Logger.Sugar().Error("GetContractSymbolOnMainNet err ", token)
		return err, ""
	}
	abiStr, err := abifile.GetAbiByToken("erc20")
	if err != nil {
		log.Logger.Sugar().Error("GetContractSymbolOnMainNet err ", token)
		return err, ""
	}
	abiJson, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		log.Logger.Sugar().Error("GetContractSymbolOnMainNet err ", token)
		return err, ""
	}
	contract := bind.NewBoundContract(common.HexToAddress(token), abiJson, client, client, client)

	var res = make([]interface{}, 0)
	err = contract.Call(nil, &res, "symbol")
	if err != nil {
		log.Logger.Sugar().Error("GetContractSymbolOnMainNet err ", token)
		return err, ""
	}
	return nil, res[0].(string)
}

// 主网
func (s *TokenSymbolService) GetMainNetTokenSymbol(token string) (error, string) {
	client, err := ethclient.Dial(config.Config.MainNet.NetUrl)
	if err != nil {
		log.Logger.Sugar().Error("GetMainNetTokenSymbol err", token, err)
		return err, ""
	}
	abiStr, err := abifile.GetAbiByToken(token)
	if err != nil {
		log.Logger.Sugar().Error("GetMainNetTokenSymbol err", token, err)
		return err, ""
	}
	abiJson, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		log.Logger.Sugar().Error("GetMainNetTokenSymbol err", token, err)
		return err, ""
	}
	contract := bind.NewBoundContract(common.HexToAddress(token), abiJson, client, client, client)
	res := make([]interface{}, 0)
	err = contract.Call(nil, &res, "symbol")
	if err != nil {
		log.Logger.Sugar().Error("GetMainNetTokenSymbol err", token, err)
		return err, ""
	}
	return nil, res[0].(string)
}

// 根据token获取远程abi文件
func (s *TokenSymbolService) GetRemoteAbiFileByToken(token string, chainId string) error {
	url := "https://api.bscscan.com/api?module=contract&action=getabi&apikey=HJ3WS4N88QJ6S7PQ8D89BD49IZIFP1JFER&address=" + token
	resbytes, err := utils.HttpGet(url, map[string]string{})
	if err != nil {
		log.Logger.Sugar().Error(err.Error())
		return err
	}
	resStr := s.FormatAbiJsonStr(string(resbytes))

	var abiJson model.AbiJson
	err = json.Unmarshal([]byte(resStr), &abiJson)
	if err != nil {
		log.Logger.Sugar().Error(err.Error())
		return err
	}
	if abiJson.Status != "1" {
		log.Logger.Sugar().Error("get remote abi file failed: status 0", err)
		return errors.New("get remote abi file failed: status 0")
	}
	abiJsonBytes, err := json.MarshalIndent(abiJson.Result, "", "\t")
	if err != nil {
		log.Logger.Sugar().Error(err.Error())
		return err
	}
	newAbiFile := abifile.GetCurrentAbPathByCaller() + "/" + token + ".abi"
	err = os.WriteFile(newAbiFile, abiJsonBytes, 0777)
	if err != nil {
		log.Logger.Sugar().Error(err.Error())
		return err
	}
	err = db.Mysql.Table("token_info").Where("chain_id=? and token=?", chainId, token).Updates(map[string]interface{}{
		"abi_file_exist": 1,
	}).Debug().Error
	if err != nil {
		log.Logger.Sugar().Error(err.Error())
		return err
	}
	return nil
}

func (s *TokenSymbolService) FormatAbiJsonStr(res string) string {
	resStr := strings.Replace(res, `\`, ``, -1)
	resStr = strings.Replace(res, `\"`, `"`, -1)
	resStr = strings.Replace(resStr, `"[{`, `[{`, -1)
	resStr = strings.Replace(resStr, `}]"`, `}]`, -1)
	return resStr
}

func (s *TokenSymbolService) CheckSymbolData(token string, chainId string, symbol string) (bool, error) {
	redisKey := "token_info:" + token + ":" + chainId
	redisTokenInfoBytes, err := db.RedisGet(redisKey)
	if len(redisTokenInfoBytes) <= 0 {
		err = s.CheckTokenInfo(token, chainId)
		if err != nil {
			log.Logger.Sugar().Error(err.Error())
			return false, err
		}
		err = db.RedisSet(redisKey, model.RedisTokenInfo{
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
		err = json.Unmarshal(redisTokenInfoBytes, &redisTokenInfo)
		if err != nil {
			log.Logger.Sugar().Error(err.Error())
			return false, err
		}
		if redisTokenInfo.Symbol == symbol {
			return false, nil
		}
		redisTokenInfo.Symbol = symbol
		err = db.RedisSet(redisKey, redisTokenInfo, 0)
		if err != nil {
			log.Logger.Sugar().Error(err.Error())
			return false, err
		}
	}
	return true, nil
}

func (s *TokenSymbolService) CheckTokenInfo(token string, chainId string) error {
	tokenInfo := model.TokenInfo{}
	err := db.Mysql.Table("token_info").Where("token=? and chain_id=?", token, chainId).First(&tokenInfo).Debug().Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tokenInfo = model.TokenInfo{}
			nowDataTime := utils.GetCurDateTimeFormat()
			tokenInfo.Token = token
			tokenInfo.ChainId = chainId
			tokenInfo.CreatedAt = nowDataTime
			tokenInfo.UpdatedAt = nowDataTime
			err = db.Mysql.Table("token_info").Create(token).Debug().Error
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}

func (s *TokenSymbolService) SaveSymbolData(token string, chainId string, symbol string) error {
	nowDataTime := utils.GetCurDateTimeFormat()
	err := db.Mysql.Table("token_info").Where("token=? and chain_id=?", token, chainId).Updates(map[string]interface{}{
		"symbol":     symbol,
		"updated_at": nowDataTime,
	}).Debug().Error
	if err != nil {
		log.Logger.Sugar().Error(err.Error())
		return err
	}
	return nil
}
