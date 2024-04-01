package smslog

import (
	"context"

	"github.com/toutmost/admin-core/api/internal/svc"
	"github.com/toutmost/admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSmsLogByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSmsLogByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSmsLogByIdLogic {
	return &GetSmsLogByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetSmsLogByIdLogic) GetSmsLogById(req *types.UUIDReq) (resp *types.SmsLogInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
