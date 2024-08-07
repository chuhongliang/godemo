package svc

import (
	"gozerodemo.com/apidemo/internal/config"
	"gozerodemo.com/apidemo/model/user"
)

type ServiceContext struct {
	Config    config.Config
	UserModel *user.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
