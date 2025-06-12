package entity

type PoolBaseInfoEntity struct {
	ChainId int `json:"chainId" binding:"required"`
}
