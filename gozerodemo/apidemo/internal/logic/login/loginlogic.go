package login

import (
	"context"
	"fmt"
	"time"

	"gozerodemo.com/apidemo/internal/custom"
	"gozerodemo.com/apidemo/internal/svc"
	"gozerodemo.com/apidemo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	userInfo, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.Username)
	fmt.Println("err", err)
	if err != nil {
		return nil, &custom.CustomError{
			Message: "用户名或密码错误",
			Code:    400,
		}
	}
	fmt.Println("userInfo", userInfo)

	return &types.LoginResp{
		Id:       1,
		Name:     "admin",
		Token:    "123456",
		ExpireAt: time.Now().Add(time.Hour * 24).Format("2006-01-02 15:04:05"),
	}, nil
}
