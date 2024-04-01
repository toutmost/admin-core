package tasklog

import (
	"context"

	"github.com/toutmost/admin-core/api/internal/svc"
	"github.com/toutmost/admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTaskLogByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTaskLogByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTaskLogByIdLogic {
	return &GetTaskLogByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetTaskLogByIdLogic) GetTaskLogById(req *types.IDReq) (resp *types.TaskLogInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
