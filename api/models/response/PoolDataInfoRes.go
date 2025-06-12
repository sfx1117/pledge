package response

type PoolDataInfoRes struct {
	Index                  int    `json:"index"`
	PoolId                 int    `json:"poolId"`
	SettleAmountLend       string `json:"settleAmountLend"`
	SettleAmountBorrow     string `json:"settleAmountBorrow"`
	FinishAmountLend       string `json:"finishAmountLend"`
	FinishAmountBorrow     string `json:"finishAmountBorrow"`
	LiquidationAmounLend   string `json:"liquidationAmounLend"`
	LiquidationAmounBorrow string `json:"liquidationAmounBorrow"`
	ChainId                string `json:"chainId"`
}
