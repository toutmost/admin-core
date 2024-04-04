package dictionary

import (
	"context"
	"github.com/toutmost/admin-common/i18n"
	"github.com/toutmost/admin-core/rpc/ent"
	"github.com/toutmost/admin-core/rpc/ent/dictionary"
	"github.com/toutmost/admin-core/rpc/ent/dictionarydetail"
	"github.com/toutmost/admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/toutmost/admin-core/rpc/internal/utils/entx"

	"github.com/toutmost/admin-core/rpc/internal/svc"
	"github.com/toutmost/admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDictionaryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteDictionaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDictionaryLogic {
	return &DeleteDictionaryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteDictionaryLogic) DeleteDictionary(in *core.IDsReq) (*core.BaseResp, error) {
	err := entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
		_, txErr := tx.DictionaryDetail.Delete().Where(dictionarydetail.HasDictionariesWith(dictionary.IDIn(in.Ids...))).Exec(l.ctx)
		if txErr != nil {
			return txErr
		}

		_, txErr = tx.Dictionary.Delete().Where(dictionary.IDIn(in.Ids...)).Exec(l.ctx)
		if txErr != nil {
			return txErr
		}

		return nil
	})
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
