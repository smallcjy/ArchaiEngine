package gate

import (
	"encoding/json"
	"game/gate/conf"
	"game/gate/server"
	reader "game/proto"
	"game/proto/systempb"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/kercylan98/vivid/src/vivid"
)

var Gateway Gate

func NewGate() *Gate {
	system := vivid.NewActorSystem(&GateConfigurator{}).StartP()
	ref := system.ActorOf(func() vivid.Actor {
		return vivid.ActorFN(Handle)
	})

	return &Gate{
		System: system,
		GateRef: ref,
	}
}

type GateConfigurator struct {}

func (g *GateConfigurator) Configure(config *vivid.ActorSystemConfig) {
	config.WithAddress(conf.ServerConfig.Addr)
}

// NewConnHandle 新连接到来消息处理函数
func Handle(ctx vivid.ActorContext) {
	message := ctx.Message()
	if msg, err := reader.ReadProto(message.([]byte), &systempb.Message{}); err == nil {
		switch msg.(*systempb.Message).MsgId {
		case systempb.MsgId_Msg_NewConn:
			Gateway.HandleNewConn(msg.(*systempb.Message))
		}
	}
}

type Gate struct {
	System vivid.ActorSystem
	GateRef vivid.ActorRef
	ConnsMap map[string]*websocket.Conn
}

func (g *Gate) HandleNewConn(message *systempb.Message) {
	client := &server.Client{}
	if err := json.Unmarshal(message.Data, client); err != nil {
		return
	}

	//生成session
	session := uuid.NewString()
	g.ConnsMap[session] = client.Extra[server.WS].(*websocket.Conn)

	// 转发消息到
}