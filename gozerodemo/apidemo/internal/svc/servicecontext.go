package svc

import (
	"context"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"gozerodemo.com/apidemo/internal/config"
	"gozerodemo.com/apidemo/model/user"
	"gozerodemo.com/greet/greet"
	"gozerodemo.com/greet/greetclient"
	"gozerodemo.com/rpcdemo/rpcdemoclient"
)

type ServiceContext struct {
	Config     config.Config
	UserModel  user.UserModel
	RpcDemoRpc rpcdemoclient.Rpcdemo
	GreetRpc   greet.GreetClient
}

func timeInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	stime := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	if err != nil {
		return err
	}

	fmt.Printf("调用 %s 方法 耗时: %v\n", method, time.Now().Sub(stime))
	return nil
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	um := user.NewUserModel(conn)

	// gr := greet.NewGreetClient(zrpc.MustNewClient(zrpc.RpcClientConf{Target: "dns:///127.0.0.1:8080"}, zrpc.WithUnaryClientInterceptor(timeInterceptor)).Conn())
	// gr := greetclient.NewGreet(zrpc.MustNewClient(zrpc.RpcClientConf{Target: "dns:///127.0.0.1:6002"}, zrpc.WithUnaryClientInterceptor(timeInterceptor)))
	// rr := rpcdemoclient.NewRpcdemo(zrpc.MustNewClient(zrpc.RpcClientConf{Target: "dns:///127.0.0.1:6001"}, zrpc.WithUnaryClientInterceptor(timeInterceptor)))

	gr := greetclient.NewGreet(zrpc.MustNewClient(c.GreetRpc, zrpc.WithUnaryClientInterceptor(timeInterceptor)))
	rr := rpcdemoclient.NewRpcdemo(zrpc.MustNewClient(c.RpcDemoRpc, zrpc.WithUnaryClientInterceptor(timeInterceptor)))

	return &ServiceContext{
		Config:     c,
		UserModel:  um,
		GreetRpc:   gr,
		RpcDemoRpc: rr,
	}
}
