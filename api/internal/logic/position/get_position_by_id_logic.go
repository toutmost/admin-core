package position

import (
	"context"
	"github.com/toutmost/admin-common/i18n"
	"github.com/toutmost/admin-core/rpc/types/core"

	"github.com/toutmost/admin-core/api/internal/svc"
	"github.com/toutmost/admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPositionByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPositionByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPositionByIdLogic {
	return &GetPositionByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetPositionByIdLogic) GetPositionById(req *types.IDReq) (resp *types.PositionInfoResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetPositionById(l.ctx, &core.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.PositionInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.PositionInfo{
			BaseIDInfo: types.BaseIDInfo{
				Id:        data.Id,
				CreatedAt: data.CreatedAt,
				UpdatedAt: data.UpdatedAt,
			},
			Status: data.Status,
			Sort:   data.Sort,
			Name:   data.Name,
			Code:   data.Code,
			Remark: data.Remark,
		},
	}, nil
}
