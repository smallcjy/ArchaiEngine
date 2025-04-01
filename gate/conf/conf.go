package conf

var (
	ServerConfig *ServerConf = &ServerConf{}
)

type ServerConf struct {
	App
	Log
	Gateway
}

type App struct {
	AppName    string
	GameId     string
	Addr       string
}

type Log struct {
	LogPath string
	IsDebug bool
	LogTime int32
}

type Gateway struct {
	InsType          string
	Disabled         bool
	ForgedInstanceId string
	AllowForged      bool
	Url              string
	UsePort          bool
	GameRouter       string
}

func (s *ServerConf) IsAws() bool {
	return s.Gateway.InsType == "AWS"
}

func (s *ServerConf) IsTencent() bool {
	return s.Gateway.InsType == "Tencent"
}