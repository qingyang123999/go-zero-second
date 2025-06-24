package logic

import (
	"api-jwt/common/jwts"
	"context"

	"api-jwt/internal/svc"
	"api-jwt/internal/types"

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

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line
	auth := l.svcCtx.Config.Auth
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		UserID:   1,
		Username: "谭沧海",
		Role:     100,
	}, auth.AccessSecret, auth.AccessExpire)
	if err != nil {
		return nil, err
	}
	return &types.LoginResponse{
		Code: 0,
		Data: token,
		Msg:  "成功",
	}, err
}
