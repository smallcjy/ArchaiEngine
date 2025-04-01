package server

// 额外字段
const (
	WS = "ws"
)

type Client struct {
	PlatformId int64                  `form:"platformid" json:"platformid,omitempty"`
	UserId     string                 `form:"userid" binding:"required" json:"userid,omitempty"`
	RoomId     string                 `form:"roomid" json:"roomid,omitempty"`
	Token      string                 `form:"token" json:"token,omitempty"`
	Robot      int                    `form:"robot" json:"robot,omitempty"`
	Extra      map[string]interface{} `form:"extra" json:"extra,omitempty"`
}