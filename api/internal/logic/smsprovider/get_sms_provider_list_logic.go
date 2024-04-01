package smsprovider

import (
	"context"

	"github.com/toutmost/admin-core/api/internal/svc"
	"github.com/toutmost/admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSmsProviderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSmsProviderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSmsProviderListLogic {
	return &GetSmsProviderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetSmsProviderListLogic) GetSmsProviderList(req *types.SmsProviderListReq) (resp *types.SmsProviderListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
