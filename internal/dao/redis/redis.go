package redis

import (
	"game/internal/conf"
	"game/internal/logs"
	"game/internal/utils/errors"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

var (
	rds    *redis.Client
	newDao *RedisDao
)

type RedisDao struct {
	Cli *redis.Client
}

func RedisStart() *errors.Error {
	con := redis.NewClient(&redis.Options{
		PoolSize: conf.GetServerConf().Redis.PoolSize, //注意连接池数量
		Addr:     conf.GetServerConf().Redis.Addr,
		Password: conf.GetServerConf().Redis.PassWord,
		DB:       conf.GetServerConf().Redis.DB,
	})
	r, err := con.Ping().Result()
	if err != nil {
		return errors.ErrRedis
	}
	logs.Info("redis init", zap.Any("conf", conf.GetServerConf().Redis), zap.String("result", r))
	rds = con
	newDao = &RedisDao{
		Cli: rds,
	}
	return nil
}

func RedisClose() {
	if rds != nil {
		rds.Close()
	}
}

func GetDao() *RedisDao {
	return newDao
}