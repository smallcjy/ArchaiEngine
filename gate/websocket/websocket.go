package ws

import (
	"encoding/json"
	"game/gate"
	"game/gate/conf"
	"game/gate/gateway"
	"game/gate/server"
	"game/logs"
	"game/proto/systempb"
	"net/http"

	"github.com/gin-gonic/gin"
	wst "github.com/gorilla/websocket"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)



var upgrader = wst.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func InitWebsocket(router *gin.Engine) {
	InitWebSocketRouter(router)
}

func InitWebSocketRouter(r *gin.Engine) {
	r.GET("/", server.HealthCheck)
	r.GET("/game/jackaroo", ServiceWebSocket)
	r.GET(conf.ServerConfig.Gateway.GameRouter, ServiceWebSocket)
	if !conf.ServerConfig.Gateway.Disabled {
		r.GET("/"+gateway.GetInstanceId()+conf.ServerConfig.Gateway.GameRouter, ServiceWebSocket)
	}
}

func ServiceWebSocket(ctx *gin.Context) {
	ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		logs.Error("Websocket Upgrade Error", zap.Error(err))
		ws.Close()
		return
	}

	client := &server.Client{}
	if err := ctx.ShouldBind(client); err != nil {
		logs.Error("Params not valid", zap.Error(err))
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	client.Extra[server.WS] = ws

	// 向gate actor发送消息有新连接到来
	data, err := json.Marshal(client)
	if err != nil {
		logs.Error("json marshal error", zap.Error(err))
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	msg, err := proto.Marshal(&systempb.Message{
		MsgId: systempb.MsgId_Msg_NewConn,
		Data: data,
	})

	if err != nil {
		logs.Error("proto marshal error", zap.Error(err))
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	gate.Gateway.System.Tell(gate.Gateway.GateRef, msg)
}