package emaillog

import (
	"context"

	"github.com/toutmost/admin-core/api/internal/svc"
	"github.com/toutmost/admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEmailLogByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetEmailLogByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEmailLogByIdLogic {
	return &GetEmailLogByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetEmailLogByIdLogic) GetEmailLogById(req *types.UUIDReq) (resp *types.EmailLogInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
