package auth

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"gozerodemo.com/apidemo/internal/svc"
	"gozerodemo.com/apidemo/internal/types"
	"gozerodemo.com/greet/greet"
	"gozerodemo.com/rpcdemo/rpcdemo"
)

type PingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PingLogic) Ping(req *types.PingReq) (resp *types.PingResp, err error) {
	reply, err := l.svcCtx.GreetRpc.Ping(l.ctx, &greet.Request{})
	if err != nil {
		return nil, err
	}
	logx.Infof("rpc ping greet result: %s", reply)

	resutl, err2 := l.svcCtx.RpcDemoRpc.Ping(l.ctx, &rpcdemo.Request{})
	if err2 != nil {
		return nil, err2
	}
	logx.Infof("rpc ping rpcdemo result: %s", resutl)
	return &types.PingResp{
		Result: "pong",
	}, nil
}
