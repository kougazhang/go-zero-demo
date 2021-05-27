package logic

import (
	"context"

	"blog/rpc/user/internal/svc"
	"blog/rpc/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteLogic) Delete(in *user.ReqUserId) (*user.CommResp, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.Model.Delete(in.Id)
	if err != nil {
		return nil, err
	}
	return &user.CommResp{Ok: true}, nil
}
