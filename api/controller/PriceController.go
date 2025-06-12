package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"pledge-backend-test/api/models/ws"
	"pledge-backend-test/log"
	"pledge-backend-test/utils"
	"strings"
	"time"
)

type PriceController struct {
}

func (p *PriceController) NewPrice(ctx *gin.Context) {
	//异常恢复机制
	defer func() {
		//使用 recover() 捕获可能的 panic
		recoverRes := recover()
		if recoverRes != nil {
			log.Logger.Sugar().Error("new price recover ", recoverRes)
		}
	}()
	//将 HTTP 连接升级为 WebSocket 连接
	coon, err := (&websocket.Upgrader{
		ReadBufferSize:   1024,            //读缓冲区大小
		WriteBufferSize:  1024,            //写缓冲区大小
		HandshakeTimeout: 5 * time.Second, //握手超时：5 秒
		CheckOrigin: func(r *http.Request) bool { //跨域检查
			return true // 允许所有跨域请求
		},
	}).Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Logger.Sugar().Error("websocket request err:", err)
		return
	}
	//生成客户端标识
	randomId := ""
	ip := ctx.RemoteIP()
	if ip == "" {
		randomId = utils.GetRandomString(32)
	} else {
		randomId = strings.Replace(ip, ".", "_", -1) + "_" + utils.GetRandomString(32)
	}
	server := &ws.Server{
		Id:       randomId,
		Socket:   coon,
		Send:     make(chan []byte, 800),
		LastTime: time.Now().Unix(),
	}
	go server.ReadAndWrite()
}
