package dao

import (
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"pledge-backend-test/api/entity"
	"pledge-backend-test/db"
)

type MultiSignDao struct {
	Id               int32  `gorm:"column:id;primaryKey"`
	SpName           string `json:"sp_name" gorm:"column:sp_name"`
	ChainId          int    `json:"chain_id" gorm:"column:chain_id"`
	SpToken          string `json:"_spToken" gorm:"column:sp_token"`
	JpName           string `json:"jp_name" gorm:"column:jp_name"`
	JpToken          string `json:"_jpToken" gorm:"column:jp_token"`
	SpAddress        string `json:"sp_address" gorm:"column:sp_address"`
	JpAddress        string `json:"jp_address" gorm:"column:jp_address"`
	SpHash           string `json:"spHash" gorm:"column:sp_hash"`
	JpHash           string `json:"jpHash" gorm:"column:jp_hash"`
	MultiSignAccount string `json:"multi_sign_account" gorm:"column:multi_sign_account"`
}

func NewMultiSignDao() *MultiSignDao {
	return &MultiSignDao{}
}

// 使用 Debug() 模式，会输出执行的 SQL 语句
func (muWs *MultiSignDao) Set(multiSign *entity.MultiSignEntity) error {
	//序列化为json格式的字节数组
	MultiSignAccountByteArr, _ := json.Marshal(multiSign.MultiSignAccount)
	//从 multi_sign 表中删除与指定 chain_id 匹配的记录
	err := db.Mysql.Table("multi_sign").Where("chain_id", multiSign.ChainId).Delete(&muWs).Debug().Error
	if err != nil {
		return errors.New("record select err" + err.Error())
	}
	err = db.Mysql.Table("multi_sign").Where("id", muWs.Id).Create(&MultiSignDao{
		SpName:           multiSign.SpName,
		ChainId:          multiSign.ChainId,
		SpToken:          multiSign.SpToken,
		SpAddress:        multiSign.SpAddress,
		SpHash:           multiSign.SpHash,
		JpName:           multiSign.JpName,
		JpToken:          multiSign.JpToken,
		JpAddress:        multiSign.JpAddress,
		JpHash:           multiSign.JpHash,
		MultiSignAccount: string(MultiSignAccountByteArr),
	}).Debug().Error
	if err != nil {
		return err
	}
	return nil
}

func (m *MultiSignDao) Get(chainId int) error {
	err := db.Mysql.Table("multi_sign").Where("chain_id", chainId).First(&m).Debug().Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		} else {
			return errors.New("record select err" + err.Error())
		}
	}
	return nil
}
