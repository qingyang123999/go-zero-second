package logic

import (
	"api-jwt/internal/svc"
	"api-jwt/internal/types"
	"context"
	"encoding/json"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (resp *types.UserInfoResponse, err error) {
	// todo: add your logic here and delete this line

	// 将login中写入的用户信息  从jwt中解析出来的信息
	userId := l.ctx.Value("user_id").(json.Number)
	fmt.Printf("%v, %T, \n", userId, userId)
	username := l.ctx.Value("username").(string)
	fmt.Printf("%v, %T, \n", username, username)
	uid, _ := userId.Int64()

	return &types.UserInfoResponse{
		Code: 0,
		Data: types.UserInfo{
			UserName: username,
			Id:       uint(uid),
			Addr:     string(l.ctx.Value("role").(json.Number)),
		},
		Msg: "成功",
	}, nil
}
