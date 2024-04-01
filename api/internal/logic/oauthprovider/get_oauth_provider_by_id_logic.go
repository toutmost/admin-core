package oauthprovider

import (
	"context"

	"github.com/toutmost/admin-core/api/internal/svc"
	"github.com/toutmost/admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOauthProviderByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOauthProviderByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOauthProviderByIdLogic {
	return &GetOauthProviderByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetOauthProviderByIdLogic) GetOauthProviderById(req *types.IDReq) (resp *types.OauthProviderInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
