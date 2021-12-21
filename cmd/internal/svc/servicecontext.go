package svc

import (
	"github.com/tal-tech/go-zero/rest"
	"go-zero-template/cmd/internal/config"
	"go-zero-template/cmd/internal/middleware"
)

type ServiceContext struct {
	Config        config.Server
	OperateRecord rest.Middleware
	Jwt           rest.Middleware
	Casbin        rest.Middleware
}

func NewServiceContext(c config.Server) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		OperateRecord: middleware.OperationRecord,
		Jwt:           middleware.JWTAuth,
		Casbin:        middleware.CasbinHandler,
	}
}
