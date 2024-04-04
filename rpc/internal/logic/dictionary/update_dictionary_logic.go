package dictionary

import (
	"context"
	"github.com/toutmost/admin-common/i18n"
	"github.com/toutmost/admin-common/utils/pointy"
	"github.com/toutmost/admin-core/rpc/internal/utils/dberrorhandler"

	"github.com/toutmost/admin-core/rpc/internal/svc"
	"github.com/toutmost/admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDictionaryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDictionaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDictionaryLogic {
	return &UpdateDictionaryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateDictionaryLogic) UpdateDictionary(in *core.DictionaryInfo) (*core.BaseResp, error) {
	err := l.svcCtx.DB.Dictionary.UpdateOneID(*in.Id).
		SetNotNilStatus(pointy.GetStatusPointer(in.Status)).
		SetNotNilTitle(in.Title).
		SetNotNilName(in.Name).
		SetNotNilDesc(in.Desc).
		Exec(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
