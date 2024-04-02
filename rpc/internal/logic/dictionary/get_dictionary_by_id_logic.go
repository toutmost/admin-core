package dictionary

import (
	"context"

	"github.com/toutmost/admin-core/rpc/internal/svc"
	"github.com/toutmost/admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDictionaryByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDictionaryByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDictionaryByIdLogic {
	return &GetDictionaryByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDictionaryByIdLogic) GetDictionaryById(in *core.IDReq) (*core.DictionaryInfo, error) {
	// todo: add your logic here and delete this line

	return &core.DictionaryInfo{}, nil
}
