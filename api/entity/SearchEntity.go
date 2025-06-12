package entity

type SearchEntity struct {
	ChainId         int    `form:"chainID" json:"chainID" binding:"required"`
	LendTokenSymobl string `form:"lend_token_symbol" json:"lend_token_symobl" binding:"omitempty"`
	State           string `form:"state" json:"state" binding:"omitempty"`
	Page            int    `form:"page" json:"page"`
	PageSize        int    `form:"pageSize" json:"pageSize"`
}
