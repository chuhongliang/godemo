package user

import (
	"context"

	"gozerodemo.com/apidemo/internal/code"
	"gozerodemo.com/apidemo/internal/custom"
	"gozerodemo.com/apidemo/internal/svc"
	"gozerodemo.com/apidemo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.GetUserInfoReq) (resp *types.GetUserInfoResp, err error) {
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return nil, &custom.CustomError{
			Code: code.USERNAME_OR_PASSWORD_ERROR,
		}
	}
	return &types.GetUserInfoResp{
		Id:        user.Id,
		Username:  user.Username,
		Avatar:    user.Avatar,
		Level:     user.Level,
		Exp:       user.Exp,
		Gold:      user.Gold,
		Diamond:   user.Diamond,
		LandLevel: user.LandLevel,
		Extra:     user.Extra,
	}, nil
}
