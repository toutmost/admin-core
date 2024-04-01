package svc

import (
	"github.com/toutmost/admin-core/api/internal/config"
	"github.com/toutmost/admin-core/api/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config: c,
	}
}
