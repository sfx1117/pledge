package kucoin

import (
	"context"
	"github.com/Kucoin/kucoin-go-sdk"
	"pledge-backend-test/db"
	"pledge-backend-test/log"
)

const ApiKeyVersion = "2"

var PledgePrice = "0.0027"
var PledgePriceChan = make(chan string, 2)

func GetExchangePrice() {
	log.Logger.Sugar().Info("GetExchangePrice")
	//从redis中获取price
	price, err := db.RedisGetString("pledge_price")
	if err != nil {
		log.Logger.Sugar().Error("get plgr price from redis err", err)
	} else {
		PledgePrice = price
	}
	//初始化 KuCoin API 服务
	service := kucoin.NewApiService(
		kucoin.ApiKeyOption("key"),
		kucoin.ApiSecretOption("secret"),
		kucoin.ApiPassPhraseOption("passphrase"),
		kucoin.ApiKeyVersionOption(ApiKeyVersion),
	)
	//获取websocket的连接令牌
	token, err := service.WebSocketPublicToken(context.Background())
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}
	//解析令牌响应
	model := &kucoin.WebSocketTokenModel{}
	err = token.ReadData(model)
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}
	//建立websocket连接
	client := service.NewWebSocketClient(model)
	connect, errors, err := client.Connect()
	if err != nil {
		log.Logger.Sugar().Error("Error:", err)
		return
	}
	//订阅价格更新频道
	ch := kucoin.NewSubscribeMessage("/market/ticker:PLGR-USDT", false)    //订阅消息通道
	uch := kucoin.NewUnsubscribeMessage("/market/ticker:PLGR-USDT", false) //取消订阅消息
	err = client.Subscribe(ch)                                             //客户端订阅消息
	if err != nil {
		log.Logger.Error(err.Error()) // Handle error
		return
	}
	//主循环
	for {
		select {
		case e := <-errors:
			client.Stop()
			_ = client.Unsubscribe(uch)
			log.Logger.Sugar().Error(e.Error())
			return
		case msg := <-connect:
			t := kucoin.TickerLevel1Model{}
			err = msg.ReadData(t)
			if err != nil {
				log.Logger.Sugar().Errorf("Failure to read: %s", err.Error())
				return
			}
			PledgePriceChan <- t.Price
			PledgePrice = t.Price
			_ = db.RedisSetString("pledge_price", PledgePrice, 0)
		}
	}
}
