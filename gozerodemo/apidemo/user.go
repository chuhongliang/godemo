package main

import (
	"flag"
	"fmt"

	"gozerodemo.com/apidemo/internal/config"
	"gozerodemo.com/apidemo/internal/handler"
	"gozerodemo.com/apidemo/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	// 需要自行将 dsn 中的 host，账号 密码配置正确
	dsn := "root:123456@tcp(127.0.0.1:3306)/gozero?charset=utf8mb4&parseTime=True&loc=Local"
	conn := sqlx.NewMysql(dsn)
	_ = conn

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)

	server.Start()
}
