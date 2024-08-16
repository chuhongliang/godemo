package logic

import (
	"context"

	"gozerodemo.com/rpcdemo/internal/svc"
	"gozerodemo.com/rpcdemo/rpcdemo"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *rpcdemo.Request) (*rpcdemo.Response, error) {
	return &rpcdemo.Response{
		Pong: "pong",
	}, nil
}
