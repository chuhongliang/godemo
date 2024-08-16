package auth

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"gozerodemo.com/apidemo/internal/code"
	"gozerodemo.com/apidemo/internal/custom"
	"gozerodemo.com/apidemo/internal/svc"
	"gozerodemo.com/apidemo/internal/types"
	"gozerodemo.com/apidemo/model/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	userInfo, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.Username)
	fmt.Println("userInfo", userInfo)
	fmt.Println("err", err)

	if userInfo != nil {
		return nil, &custom.CustomError{
			Code: code.USERNAME_ALREADY_EXISTS,
		}
	}
	now := sql.NullTime{Time: time.Now(), Valid: true}
	l.svcCtx.UserModel.Insert(l.ctx, &user.User{
		Username:  req.Username,
		Password:  req.Password,
		Level:     1,
		Exp:       0,
		Gold:      0,
		Diamond:   0,
		LandLevel: 1,
		LoginAt:   now,
		Extra:     "{}",
	})

	return &types.RegisterResp{
		Id:       1,
		Name:     "admin",
		Token:    "123456",
		ExpireAt: time.Now().Add(time.Hour * 24).Format("2006-01-02 15:04:05"),
	}, nil
}
