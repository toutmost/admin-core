package config

import (
	"github.com/toutmost/admin-common/config"
	"github.com/toutmost/admin-common/plugins/casbin"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DatabaseConf config.DatabaseConf
	CasbinConf   casbin.CasbinConf
	RedisConf    config.RedisConf
}
