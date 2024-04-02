package oauthprovider

import (
	"context"

	"github.com/toutmost/admin-core/rpc/internal/svc"
	"github.com/toutmost/admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOauthProviderByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOauthProviderByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOauthProviderByIdLogic {
	return &GetOauthProviderByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOauthProviderByIdLogic) GetOauthProviderById(in *core.IDReq) (*core.OauthProviderInfo, error) {
	// todo: add your logic here and delete this line

	return &core.OauthProviderInfo{}, nil
}
