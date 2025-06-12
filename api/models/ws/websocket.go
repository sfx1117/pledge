package ws

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"pledge-backend-test/api/models/kucoin"
	"pledge-backend-test/config"
	"pledge-backend-test/log"
	"sync"
	"time"
)

const SuccessCode = 0
const PongCode = 1
const ErrorCode = -1

type Server struct {
	sync.Mutex //提供对结构体的并发访问控制,保护对 WebSocket 连接和共享数据的并发访问,在多 goroutine 环境下确保线程安全
	Id         string
	Socket     *websocket.Conn
	Send       chan []byte
	LastTime   int64
}

type ServerManager struct {
	Servers    sync.Map
	Brandcast  chan []byte
	Register   chan *Server
	UnRegister chan *Server
}

type Message struct {
	Code int    `json:"code"`
	Data string `json:"data"`
}

var Manager = ServerManager{}
var UserPingPongDurTime = config.Config.Env.WssTimeOutDuration

func (s *Server) SendToClient(data string, code int) {
	//加锁
	s.Lock()
	defer s.Unlock()
	//将消息转换为byte
	messageBytes, err := json.Marshal(Message{
		Code: code,
		Data: data,
	})
	//写消息
	err = s.Socket.WriteMessage(websocket.TextMessage, messageBytes)

	if err != nil {
		log.Logger.Sugar().Error(s.Id+" SendToClient err ", err)
	}
}

func (s *Server) ReadAndWrite() {
	errChan := make(chan error)
	//注册连接
	Manager.Servers.Store(s.Id, s)
	//释放资源
	defer func() {
		Manager.Servers.Delete(s)
		_ = s.Socket.Close()
		close(s.Send)
	}()
	//write
	go func() {
		for {
			select {
			case msg, ok := <-s.Send:
				if !ok {
					errChan <- errors.New("write message err")
					return
				}
				s.SendToClient(string(msg), SuccessCode)
			}
		}
	}()
	//read
	go func() {
		for {
			_, msg, err := s.Socket.ReadMessage()
			if err != nil {
				log.Logger.Sugar().Error(err.Error())
				errChan <- err
				return
			}
			//心跳
			if string(msg) == "ping" || string(msg) == `"ping"` || string(msg) == `'ping'` {
				s.LastTime = time.Now().Unix()
				s.SendToClient("pong", PongCode)
			}
			continue
		}
	}()
	for {
		select {
		//检查超时
		case <-time.After(time.Second):
			if time.Now().Unix()-s.LastTime > UserPingPongDurTime {
				s.SendToClient("", ErrorCode)
				return
			}
			//错误处理
		case err := <-errChan:
			log.Logger.Sugar().Error(s.Id, " ReadAndWrite returned ", err)
			return
		}
	}
}

func StartServer() {
	log.Logger.Sugar().Info("WsServer start")
	for {
		select {
		case price, ok := <-kucoin.PledgePriceChan:
			if ok {
				Manager.Servers.Range(func(key, value any) bool {
					value.(*Server).SendToClient(price, SuccessCode)
					return true
				})
			}
		}
	}
}
