package entity

type MultiSignEntity struct {
	SpName           string   `json:"sp_name" binding:"required"`
	ChainId          int      `json:"chain_id"`
	SpToken          string   `json:"sp_token"`
	SpAddress        string   `json:"sp_address"`
	SpHash           string   `json:"sp_hash"`
	JpName           string   `json:"jp_name"`
	JpToken          string   `json:"jp_token"`
	JpAddress        string   `json:"jp_address"`
	JpHash           string   `json:"jp_hash"`
	MultiSignAccount []string `json:"multi_sign_account"`
}

type GetMultiSign struct {
	ChainId int `json:"chain_id"`
}
