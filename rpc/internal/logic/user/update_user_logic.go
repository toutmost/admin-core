package user

import (
	"context"
	"github.com/toutmost/admin-common/enum/common"
	"github.com/toutmost/admin-common/i18n"
	"github.com/toutmost/admin-common/utils/encrypt"
	"github.com/toutmost/admin-common/utils/pointy"
	"github.com/toutmost/admin-common/utils/uuidx"
	"github.com/toutmost/admin-core/rpc/internal/logic/token"
	"github.com/toutmost/admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/toutmost/admin-core/rpc/internal/utils/entx"

	"github.com/toutmost/admin-core/rpc/ent"

	"github.com/toutmost/admin-core/rpc/internal/svc"
	"github.com/toutmost/admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *core.UserInfo) (*core.BaseResp, error) {
	err := entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
		updateQuery := tx.User.UpdateOneID(uuidx.ParseUUIDString(*in.Id)).
			SetNotNilUsername(in.Username).
			SetNotNilNickname(in.Nickname).
			SetNotNilEmail(in.Email).
			SetNotNilMobile(in.Mobile).
			SetNotNilAvatar(in.Avatar).
			SetNotNilHomePath(in.HomePath).
			SetNotNilDescription(in.Description).
			SetNotNilDepartmentID(in.DepartmentId).
			SetNotNilStatus(pointy.GetStatusPointer(in.Status))

		if in.Password != nil {
			updateQuery = updateQuery.SetNotNilPassword(pointy.GetPointer(encrypt.BcryptEncrypt(*in.Password)))
		}

		if in.RoleIds != nil {
			err := tx.User.UpdateOneID(uuidx.ParseUUIDString(*in.Id)).ClearRoles().Exec(l.ctx)
			if err != nil {
				return err
			}

			updateQuery = updateQuery.AddRoleIDs(in.RoleIds...)
		}

		if in.PositionIds != nil {
			err := tx.User.UpdateOneID(uuidx.ParseUUIDString(*in.Id)).ClearPositions().Exec(l.ctx)
			if err != nil {
				return err
			}

			updateQuery = updateQuery.AddPositionIDs(in.PositionIds...)
		}

		if in.Password != nil || in.RoleIds != nil || in.PositionIds != nil || (in.Status != nil && uint8(*in.Status) != common.StatusNormal) {
			_, err := token.NewBlockUserAllTokenLogic(l.ctx, l.svcCtx).BlockUserAllToken(&core.UUIDReq{Id: *in.Id})
			if err != nil {
				return err
			}
		}

		return updateQuery.Exec(l.ctx)
	})
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseResp{
		Msg: i18n.Success,
	}, nil
}
