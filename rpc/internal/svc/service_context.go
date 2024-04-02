package svc

import (
	"github.com/redis/go-redis/v9"
	"github.com/toutmost/admin-core/rpc/ent"
	"github.com/toutmost/admin-core/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/logx"

	_ "github.com/toutmost/admin-core/rpc/ent/runtime"
)

type ServiceContext struct {
	Config config.Config
	DB     *ent.Client
	Redis  redis.UniversalClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := ent.NewClient(
		ent.Log(logx.Info), // logger
		ent.Driver(c.DatabaseConf.NewNoCacheDriver()),
		ent.Debug(), // debug mode
	)

	return &ServiceContext{
		Config: c,
		DB:     db,
		Redis:  c.RedisConf.MustNewUniversalRedis(),
	}
}
