package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gozerodemo.com/apidemo/internal/config"
	"gozerodemo.com/apidemo/model/user"
)

type ServiceContext struct {
	Config    config.Config
	UserModel user.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	um := user.NewUserModel(conn)

	return &ServiceContext{
		Config:    c,
		UserModel: um,
	}
}
