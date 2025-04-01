package main

import (
	reader "game/proto"
	"game/proto/systempb"

	"google.golang.org/protobuf/proto"
)

func main() {
	msg := &systempb.Message{
		MsgId: systempb.MsgId_Msg_None,
		Data: []byte("hello"),
	}
	data, _ := proto.Marshal(msg)
	res, _ := reader.ReadProto(data, &systempb.Message{})
	println(res.(*systempb.Message).MsgId)
}