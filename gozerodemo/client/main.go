package main

import (
	"context"
	"log"

	"github.com/zeromicro/go-zero/zrpc"
	"gozerodemo.com/greet/greet"
)

func main() {
	conn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Target: "dns:///127.0.0.1:8080",
	})
	client := greet.NewGreetClient(conn.Conn())
	resp, err := client.Ping(context.Background(), &greet.Request{})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(resp)
}
