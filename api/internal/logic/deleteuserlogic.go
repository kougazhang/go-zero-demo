package logic

import (
	"blog/rpc/user/users"
	"context"

	"blog/api/internal/svc"
	"blog/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteUserLogic {
	return DeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserLogic) DeleteUser(req types.ReqUserId) (*types.CommResp, error) {
	// todo: add your logic here and delete this line
	resp, err := l.svcCtx.User.Delete(l.ctx, &users.ReqUserId{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.CommResp{Ok: resp.Ok}, nil
}
