package logic

import (
	"context"

	"blog/rpc/user/internal/svc"
	"blog/rpc/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetAllLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllLogic {
	return &GetAllLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllLogic) GetAll(in *user.ReqGetAll) (*user.RespGetAll, error) {
	// todo: add your logic here and delete this line
	all, err := l.svcCtx.Model.FindAll()
	if err != nil {
		return nil, err
	}
	res := make([]*user.User, len(all))
	for i := 0; i < len(all); i++ {
		item := all[i]
		res[i] = &user.User{
			Id:       item.Id,
			Username: item.Username,
			Password: item.Password,
		}
	}
	return &user.RespGetAll{Users: res}, nil
}
