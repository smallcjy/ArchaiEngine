package conf

import (
	"game/internal/utils/errors"
	"github.com/spf13/viper"
)

var (
	ConfigPath string
	serverConfig *ServerConf = &ServerConf{}
)

type option func(*ServerConf)

func InitConf(opts ...option) *errors.Error{
	if ConfigPath == "" {
		return errors.ErrConf
	}
	v := viper.New()
	v.SetConfigFile(ConfigPath)
	if err := v.ReadInConfig(); err != nil {
		return errors.ErrConf
	}
	if err := v.Unmarshal(serverConfig); err != nil {
		return errors.ErrConf
	}

	for _, opt := range opts {
		opt(serverConfig)
	}
	return nil
}

type ServerConf struct {
	App
	Log
	Server
	Mysql
	Redis
	Gateway
	Nats
	Rpc
}

type App struct {
	AppName    string
	AppSecret  string
	GameId     int32
	ListenAddr int32
	RunMode    string
	SdkAddr    string
	IsWeb      bool
	Robot      bool
	SignSecret string
}

type Log struct {
	LogPath string
	IsDebug bool
	LogTime int32
}

type Server struct {
	MsgBuffSize  int64
	SeqNum       int64
	ProxyNodeNum int64
}

type Mysql struct {
	Addr     string
	UserName string
	PassWord string
	DataBase string
}
type Redis struct {
	PoolSize int
	Addr     string
	PassWord string
	DB       int
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

type Nats struct {
	NatsAddr     string
	NatsUser     string
	NatsPassword string
}

type Rpc struct {
	RankService string
}

func GetServerConf() *ServerConf {
	return serverConfig
}


func (s *ServerConf) IsAws() bool {
	return s.Gateway.InsType == "AWS"
}

func (s *ServerConf) IsTencent() bool {
	return s.Gateway.InsType == "Tencent"
}
