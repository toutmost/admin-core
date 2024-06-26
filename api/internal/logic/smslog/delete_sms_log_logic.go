package smslog

import (
	"context"

	"github.com/toutmost/admin-core/api/internal/svc"
	"github.com/toutmost/admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSmsLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSmsLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSmsLogLogic {
	return &DeleteSmsLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *DeleteSmsLogLogic) DeleteSmsLog(req *types.UUIDsReq) (resp *types.BaseMsgResp, err error) {
	// todo: add your logic here and delete this line

	return
}
