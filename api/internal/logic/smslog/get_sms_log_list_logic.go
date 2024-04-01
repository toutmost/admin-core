package smslog

import (
	"context"

	"github.com/toutmost/admin-core/api/internal/svc"
	"github.com/toutmost/admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSmsLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSmsLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSmsLogListLogic {
	return &GetSmsLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetSmsLogListLogic) GetSmsLogList(req *types.SmsLogListReq) (resp *types.SmsLogListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
