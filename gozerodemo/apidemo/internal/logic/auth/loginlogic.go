package auth

import (
	"context"
	"fmt"
	"time"

	"gozerodemo.com/apidemo/internal/code"
	"gozerodemo.com/apidemo/internal/custom"
	"gozerodemo.com/apidemo/internal/svc"
	"gozerodemo.com/apidemo/internal/types"

	"github.com/dgrijalva/jwt-go"
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
			Code: code.USERNAME_OR_PASSWORD_ERROR,
		}
	}
	if userInfo.Password != req.Password {
		return nil, &custom.CustomError{
			Code: code.USERNAME_OR_PASSWORD_ERROR,
		}
	}
	fmt.Println("userInfo", userInfo)
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, userInfo.Id)
	if err != nil {
		return nil, &custom.CustomError{
			Code: code.CREATE_TOKEN_ERROR,
		}
	}
	return &types.LoginResp{
		Id:       1,
		Name:     "admin",
		Token:    jwtToken,
		ExpireAt: time.Now().Add(time.Hour * 24).Format("2006-01-02 15:04:05"),
	}, nil
}

// @secretKey: JWT 加解密密钥
// @iat: 时间戳
// @seconds: 过期时间，单位秒
// @payload: 数据载体
func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds int64, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
