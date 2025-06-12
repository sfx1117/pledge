package service

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"pledge-backend-test/config"
	"pledge-backend-test/contract/bindings"
	"pledge-backend-test/db"
	"pledge-backend-test/log"
	"pledge-backend-test/schedule/model"
	"pledge-backend-test/utils"
	"strings"
)

type PoolService struct {
}

func NewPoolService() *PoolService {
	return &PoolService{}
}

func (s *PoolService) UpdateAllPoolInfo() {

}
func (s *PoolService) UpdataAllPoolInfo() {
	s.UpdatePoolInfo(config.Config.TestNet.PlgrAddress, config.Config.TestNet.NetUrl, config.Config.TestNet.ChainId)
	s.UpdatePoolInfo(config.Config.MainNet.PlgrAddress, config.Config.MainNet.NetUrl, config.Config.MainNet.ChainId)
}

func (s *PoolService) UpdatePoolInfo(contractAddress, network, chainId string) {
	log.Logger.Sugar().Info("UpdatePoolInfo", contractAddress+" "+network)
	client, err := ethclient.Dial(network)
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}
	pledgePoolToken, err := bindings.NewPledgePoolToken(common.HexToAddress(contractAddress), client)
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}
	//借入（我向别人借了多少钱）
	borrowFee, err := pledgePoolToken.PledgePoolTokenCaller.BorrowFee(nil)
	//借出
	lendFee, err := pledgePoolToken.PledgePoolTokenCaller.LendFee(nil)
	poolLength, err := pledgePoolToken.PledgePoolTokenCaller.PoolLength(nil)
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}
	for i := 0; i < int(poolLength.Int64()); i++ {
		log.Logger.Sugar().Info("UpdatePoolInfo", i)
		poolId := utils.IntToString(i + 1)
		//1、获取PoolBaseInfo
		baseInfo, err := pledgePoolToken.PledgePoolTokenCaller.PoolBaseInfo(nil, big.NewInt(int64(i)))
		if err != nil {
			log.Logger.Sugar().Info("UpdatePoolInfo PoolBaseInfo err", poolId, err)
			continue
		}
		_, borrowTokenInfo := model.NewTokenInfo().GetTokenInfo(baseInfo.BorrowToken.String(), chainId)
		_, lendTokenInfo := model.NewTokenInfo().GetTokenInfo(baseInfo.LendToken.String(), chainId)

		borrowTokenJson, _ := json.Marshal(model.BorrowToken{
			BorrowFee:  borrowFee.String(),
			TokenLogo:  borrowTokenInfo.Logo,
			TokenName:  borrowTokenInfo.Symbol,
			TokenPrice: borrowTokenInfo.Price,
		})
		lendTokenJson, _ := json.Marshal(model.LendToken{
			LendFee:    lendFee.String(),
			TokenLogo:  lendTokenInfo.Logo,
			TokenName:  lendTokenInfo.Symbol,
			TokenPrice: lendTokenInfo.Price,
		})
		poolBase := model.PoolBase{
			SettleTime:             baseInfo.SettleTime.String(),
			PoolId:                 utils.StringToInt(poolId),
			ChainId:                chainId,
			EndTime:                baseInfo.EndTime.String(),
			InterestRate:           baseInfo.InterestRate.String(),
			MaxSupply:              baseInfo.MaxSupply.String(),
			LendSupply:             baseInfo.LendSupply.String(),
			BorrowSupply:           baseInfo.BorrowSupply.String(),
			MartgageRate:           baseInfo.MartgageRate.String(),
			LendToken:              baseInfo.LendToken.String(),
			LendTokenSymbol:        lendTokenInfo.Symbol,
			LendTokenInfo:          string(lendTokenJson),
			BorrowToken:            baseInfo.BorrowToken.String(),
			BorrowTokenSymbol:      borrowTokenInfo.Symbol,
			BorrowTokenInfo:        string(borrowTokenJson),
			State:                  utils.IntToString(int(baseInfo.State)),
			SpCoin:                 baseInfo.SpCoin.String(),
			JpCoin:                 baseInfo.JpCoin.String(),
			AutoLiquidateThreshold: baseInfo.AutoLiquidateThreshold.String(),
		}
		//poolBase
		pbRedisKey := "base_info:pool_" + chainId + "_" + poolId
		hasInfoData, byteBaseInfoStr, baseInfoMd5Str := s.GetPoolMd5(&poolBase, pbRedisKey)
		//redis中没有  新数据
		if !hasInfoData || (byteBaseInfoStr != baseInfoMd5Str) {
			err := model.NewPoolBase().SavePoolBase(chainId, poolId, &poolBase)
			if err != nil {
				log.Logger.Sugar().Error("SavePoolBase err ", chainId, poolId)
			}
			_ = db.RedisSet(pbRedisKey, baseInfoMd5Str, 60*30)
		}
		//2、获取poolData
		poolDataInfo, err := pledgePoolToken.PledgePoolTokenCaller.PoolDataInfo(nil, big.NewInt(int64(i)))
		if err != nil {
			log.Logger.Sugar().Info("UpdatePoolInfo PoolDataInfo err", poolId, err)
			continue
		}
		pdRedisKey := "data_info:pool_" + chainId + "_" + poolId
		hasInfoData, byteDataInfoStr, dataInfoMd5Str := s.GetPoolMd5(&poolBase, pdRedisKey)
		//redis中没有  则为新数据
		if !hasInfoData || (byteDataInfoStr != dataInfoMd5Str) {
			poolData := model.PoolData{
				PoolId:                 poolId,
				ChainId:                chainId,
				FinishAmountBorrow:     poolDataInfo.FinishAmountBorrow.String(),
				FinishAmountLend:       poolDataInfo.FinishAmountLend.String(),
				LiquidationAmounBorrow: poolDataInfo.LiquidationAmounBorrow.String(),
				LiquidationAmounLend:   poolDataInfo.LiquidationAmounLend.String(),
				SettleAmountBorrow:     poolDataInfo.SettleAmountBorrow.String(),
				SettleAmountLend:       poolDataInfo.SettleAmountLend.String(),
			}
			err := model.NewPoolData().SavePoolData(chainId, poolId, &poolData)
			if err != nil {
				log.Logger.Sugar().Error("SavePoolBase err ", chainId, poolId)
			}
			_ = db.RedisSet(pdRedisKey, byteDataInfoStr, 60*30)
		}
	}
}

func (s *PoolService) GetPoolMd5(poolBase *model.PoolBase, redisKey string) (bool, string, string) {
	poolBaseBytes, _ := json.Marshal(poolBase)
	poolBaseMd5Str := utils.Md5(string(poolBaseBytes))
	redisInfoBytes, _ := db.RedisGet(redisKey)
	if len(redisInfoBytes) > 0 {
		return true, strings.Trim(string(redisInfoBytes), `"'`), poolBaseMd5Str
	} else {
		return false, strings.Trim(string(redisInfoBytes), `"'`), poolBaseMd5Str
	}
}
