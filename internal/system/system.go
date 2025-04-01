package system

import (
	"game/internal/conf"
	"game/internal/dao/mysql"
	"game/internal/dao/redis"
	"game/internal/rpc"
	"game/internal/utils/errors"
)

func NewSystem(configPath string) (*System, *errors.Error) {
	// Load Config
	conf.ConfigPath = configPath
	if err := conf.InitConf(); err != nil {
		return nil, err
	}
	// Start Dao
	if err := mysql.MysqlStart(); err != nil {
		return nil, err
	}
	// Start Redis
	if err := redis.RedisStart(); err != nil {
		return nil, err
	}
	// Start RpcService
	rpc.InitRankService(conf.GetServerConf().RankService)

	return &System{
		Conf: conf.GetServerConf(),
		MysqlDao: mysql.GetDao(),
		RedisDao: redis.GetDao(),
	}, nil
}

type System struct {
	Conf *conf.ServerConf
	MysqlDao *mysql.MysqlDao
	RedisDao *redis.RedisDao
}