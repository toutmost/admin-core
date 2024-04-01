package emailprovider

import (
	"context"

	"github.com/toutmost/admin-core/api/internal/svc"
	"github.com/toutmost/admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEmailProviderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetEmailProviderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEmailProviderListLogic {
	return &GetEmailProviderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetEmailProviderListLogic) GetEmailProviderList(req *types.EmailProviderListReq) (resp *types.EmailProviderListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
