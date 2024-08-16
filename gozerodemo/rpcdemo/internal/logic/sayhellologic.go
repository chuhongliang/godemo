package logic

import (
	"context"

	"gozerodemo.com/rpcdemo/internal/svc"
	"gozerodemo.com/rpcdemo/rpcdemo"

	"github.com/zeromicro/go-zero/core/logx"
)

type SayHelloLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSayHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SayHelloLogic {
	return &SayHelloLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SayHelloLogic) SayHello(in *rpcdemo.SayHelloReq) (*rpcdemo.SayHelloResp, error) {
	// todo: add your logic here and delete this line

	return &rpcdemo.SayHelloResp{}, nil
}
