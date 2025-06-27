package logic

import (
	"context"
	"errors"

	"api-gorm/internal/svc"
	"api-gorm/internal/types"
	"api-gorm/model"

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

	var user model.User
	err = l.svcCtx.DB.Take(&user, "username = ? and password = ?", req.UserName, req.Password).Error
	if err != nil {
		return nil, errors.New("查询user失败")
	}
	return &types.LoginResponse{
		Msg: user.Username,
	}, nil
}
