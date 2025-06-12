package model

import (
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"pledge-backend-test/db"
)

type TokenInfo struct {
	Id           int    `gorm:"column:id;primaryKey"`
	Logo         string `json:"logo" gorm:"column:logo"`
	Token        string `json:"token" gorm:"column:token"`
	Symbol       string `json:"symbol" gorm:"column:symbol"`
	ChainId      string `json:"chain_id" gorm:"column:chain_id"`
	Price        string `json:"price" gorm:"column:price"`
	Decimals     int    `json:"decimals" gorm:"column:decimals"`
	AbiFileExist int    `json:"abi_file_exist" gorm:"column:abi_file_exist"`
	CreatedAt    string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt    string `json:"updated_at" gorm:"column:updated_at"`
}

func NewTokenInfo() *TokenInfo {
	return &TokenInfo{}
}

// 获取tokenInfo
func (t *TokenInfo) GetTokenInfo(token string, chainId string) (error, TokenInfo) {
	var tokenInfo TokenInfo
	redisKey := "tokenInfo:" + chainId + ":" + token
	tokenInfoBytes, _ := db.RedisGet(redisKey)
	//redis中没有  查询数据库
	if len(tokenInfoBytes) < 0 {
		err := db.Mysql.Table("token_info").Where("token=? and chain_id=?", token, chainId).First(&tokenInfo).Debug().Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, tokenInfo
			} else {
				return errors.New("record select err " + err.Error()), tokenInfo
			}
		}
		_ = db.RedisSet(redisKey, RedisTokenInfo{
			Token:   token,
			ChainId: chainId,
			Logo:    tokenInfo.Logo,
			Symbol:  tokenInfo.Symbol,
			Price:   tokenInfo.Price,
		}, 0)
		return nil, tokenInfo
	} else {
		redisTokenInfo := RedisTokenInfo{}
		err := json.Unmarshal(tokenInfoBytes, redisTokenInfo)
		if err != nil {
			return errors.New("record Unmarshal err " + err.Error()), tokenInfo
		}
		return nil, TokenInfo{
			Logo:    redisTokenInfo.Logo,
			Token:   redisTokenInfo.Token,
			ChainId: redisTokenInfo.ChainId,
			Symbol:  redisTokenInfo.Symbol,
			Price:   redisTokenInfo.Price,
		}
	}
}
