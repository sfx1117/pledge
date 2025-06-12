package response

import "time"

type TokenListRes struct {
	Name      string    `json:"name"`
	LogoURI   string    `json:"logoURI"`
	Tokens    []Token   `json:"tokens"`
	Version   Version   `json:"version"`
	Timestamp time.Time `json:"timestamp"`
}

// omitempty  当字段为空时，json会忽略该字段
type Token struct {
	Name     string `json:"name"`
	Decimals int    `json:"decimals"`
	Symbol   string `json:"symbol"`
	Address  string `json:"address"`
	ChainId  int    `json:"chainId"`
	LogoURI  string `json:"LogoURI,omitempty"`
}
type Version struct {
	Major int `json:"major"`
	Minor int `json:"minor"`
	Patch int `json:"patch"`
}

type TokenInfo struct {
	Id      int32  `json:"-"`
	Symbol  string `json:"symbol"`
	Token   string `json:"token"`
	ChainId int    `json:"chain_id"`
}
