package user

import (
	"context"

	"github.com/toutmost/admin-core/api/internal/svc"
	"github.com/toutmost/admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *CreateUserLogic) CreateUser(req *types.UserInfo) (resp *types.BaseMsgResp, err error) {
	// todo: add your logic here and delete this line

	return
}
