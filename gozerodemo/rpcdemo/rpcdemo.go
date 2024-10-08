package main

import (
	"flag"
	"fmt"

	"gozerodemo.com/rpcdemo/internal/config"
	"gozerodemo.com/rpcdemo/internal/server"
	"gozerodemo.com/rpcdemo/internal/svc"
	"gozerodemo.com/rpcdemo/rpcdemo"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/rpcdemo.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		rpcdemo.RegisterRpcdemoServer(grpcServer, server.NewRpcdemoServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

// grpcurl -plaintext 127.0.0.1:8080 rpcdemo.Rpcdemo/Ping
// grpcurl -plaintext 127.0.0.1:8080 rpcdemo.Rpcdemo/SayHello
