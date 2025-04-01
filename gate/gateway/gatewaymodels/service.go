package gatewaymodels

// ServiceStatus 服务状态
type ServiceStatus struct {
	ConnNum        int    `json:"conn_num"`         // 当前连接数
	RoomNum        int    `json:"room_num"`         // 当前房间数
	PlayingRoomNum int    `json:"playing_room_num"` // 当前游戏中的房间数
	Unbind         bool   `json:"unbind"`           // 是否为解绑状态
	Ok             bool   `json:"ok"`               // 服务是否正常
	Forged         bool   `json:"forged"`           // 是否为伪造的非腾讯云实例
	GameRouter     string `json:"game_router"`      // 游戏路由
	UsePort        bool   `json:"use_port"`         // 是否需要携带端口访问
}

// Service 服务数据模型
type Service struct {
	Id            string          `json:"id"`      // 被用于路由的服务 ID
	Port          int64           `json:"port"`    // 端口
	Version       int             `json:"version"` // 版本
	ServiceStatus `json:"status"` // 服务状态

	notifiedExit bool // 是否已通知退出
}

// LoadScore 计算当前负载分数，当前负载分数越高，负载越高，负载分数范围为 0.0 ~ 1.0
func (s *Service) LoadScore() float64 {
	// 设置阈值和权重
	playingConnNumWeight := 0.6
	connNumWeight := 0.2
	roomNumWeight := 0.1
	playingRoomNumWeight := 0.1

	// 计算每个指标的得分
	connNumScore := float64(s.ConnNum) / 100.0
	roomNumScore := float64(s.RoomNum) / 50.0
	playingRoomNumScore := float64(s.PlayingRoomNum) / 20.0

	// 计算综合负载分数
	return playingConnNumWeight +
		connNumScore*connNumWeight +
		roomNumScore*roomNumWeight +
		playingRoomNumScore*playingRoomNumWeight
}
