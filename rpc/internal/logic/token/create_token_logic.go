package token

import (
	"context"
	"github.com/toutmost/admin-common/i18n"
	"github.com/toutmost/admin-common/utils/pointy"
	"github.com/toutmost/admin-common/utils/uuidx"
	"github.com/toutmost/admin-core/rpc/internal/utils/dberrorhandler"

	"github.com/toutmost/admin-core/rpc/internal/svc"
	"github.com/toutmost/admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTokenLogic {
	return &CreateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Token management
func (l *CreateTokenLogic) CreateToken(in *core.TokenInfo) (*core.BaseUUIDResp, error) {
	result, err := l.svcCtx.DB.Token.Create().
		SetNotNilStatus(pointy.GetStatusPointer(in.Status)).
		SetNotNilUUID(uuidx.ParseUUIDStringToPointer(in.Uuid)).
		SetNotNilToken(in.Token).
		SetNotNilSource(in.Source).
		SetNotNilUsername(in.Username).
		SetNotNilExpiredAt(pointy.GetTimeMilliPointer(in.ExpiredAt)).
		Save(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseUUIDResp{Id: result.ID.String(), Msg: i18n.CreateSuccess}, nil
}
