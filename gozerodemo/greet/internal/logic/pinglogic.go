package logic

import (
	"context"

	"gozerodemo.com/greet/greet"
	"gozerodemo.com/greet/internal/svc"

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

func (l *PingLogic) Ping(in *greet.Request) (*greet.Response, error) {
	return &greet.Response{
		Pong: "pong",
	}, nil
}
