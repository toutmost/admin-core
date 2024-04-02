package oauthprovider

import (
	"context"

	"github.com/toutmost/admin-core/rpc/internal/svc"
	"github.com/toutmost/admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOauthProviderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOauthProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOauthProviderLogic {
	return &CreateOauthProviderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// OauthProvider management
func (l *CreateOauthProviderLogic) CreateOauthProvider(in *core.OauthProviderInfo) (*core.BaseIDResp, error) {
	// todo: add your logic here and delete this line

	return &core.BaseIDResp{}, nil
}
