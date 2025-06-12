package response

type PoolBaseInfoRes struct {
	Index                  int             `json:"index"`
	PoolId                 int             `json:"poolId"`
	SettleTime             string          `json:"settleTime"`
	InterestRate           string          `json:"interestRate"`
	MaxSupply              string          `json:"maxSupply"`
	LendSupply             string          `json:"lendSupply"`
	BorrowSupply           string          `json:"borrowSupply"`
	MartgageRate           string          `json:"martgageRate"`
	LendToken              string          `json:"lendToken"`
	BorrowToken            string          `json:"borrowToken"`
	State                  string          `json:"state"`
	JpCoin                 string          `json:"jpCoin"`
	SpCoin                 string          `json:"spCoin"`
	AutoLiquidateThreshold string          `json:"autoLiquidateThreshold"`
	BorrowTokenInfo        BorrowTokenInfo `json:"borrowTokenInfo"`
	LendTokenInfo          LendTokenInfo   `json:"lendTokenInfo"`
	ChainId                string          `json:"chainId"`
	LendTokenSymbol        string          `json:"lendTokenSymbol"`
	BorrowTokenSymbol      string          `json:"borrowTokenSymbol"`
}
type BorrowTokenInfo struct {
	BorrowFee  string `json:"borrowFee"`
	TokenLogo  string `json:"tokenLogo"`
	TokenName  string `json:"tokenName"`
	TokenPrice string `json:"tokenPrice"`
}
type LendTokenInfo struct {
	LendFee    string `json:"lendFee"`
	TokenLogo  string `json:"tokenLogo"`
	TokenName  string `json:"tokenName"`
	TokenPrice string `json:"tokenPrice"`
}
