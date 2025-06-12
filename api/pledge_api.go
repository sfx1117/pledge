package main

import (
	"github.com/gin-gonic/gin"
	"pledge-backend-test/api/middleware"
	"pledge-backend-test/api/models/kucoin"
	"pledge-backend-test/api/models/ws"
	"pledge-backend-test/api/route"
	"pledge-backend-test/api/static"
	"pledge-backend-test/api/validate"
	"pledge-backend-test/config"
	"pledge-backend-test/db"
)

func main() {
	//初始化mysql
	db.InitMysql()
	//初始化redis
	db.InitReids()
	//gin bind go-playground-validator
	validate.BindingValidator()
	//启动websocket server
	go ws.StartServer()
	//get plgr price from kucoin-exchange
	go kucoin.GetExchangePrice()

	//gin start
	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()
	staticPath := static.GetCurrentAbPathByCaller()
	app.Static("/static/", staticPath)
	app.Use(middleware.Cors()) // 「 Cross domain Middleware 」
	route.InitRoute(app)
	_ = app.Run(":" + config.Config.Env.Port)
}
