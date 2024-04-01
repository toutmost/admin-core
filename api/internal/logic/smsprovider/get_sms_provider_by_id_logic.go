package smsprovider

import (
	"context"

	"github.com/toutmost/admin-core/api/internal/svc"
	"github.com/toutmost/admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSmsProviderByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSmsProviderByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSmsProviderByIdLogic {
	return &GetSmsProviderByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetSmsProviderByIdLogic) GetSmsProviderById(req *types.IDReq) (resp *types.SmsProviderInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
