package authority

import (
	"context"
	"github.com/toutmost/admin-common/i18n"
	"github.com/toutmost/admin-core/rpc/types/core"

	"github.com/toutmost/admin-core/api/internal/svc"
	"github.com/toutmost/admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApiAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetApiAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApiAuthorityLogic {
	return &GetApiAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetApiAuthorityLogic) GetApiAuthority(req *types.IDReq) (resp *types.ApiAuthorityListResp, err error) {
	roleData, err := l.svcCtx.CoreRpc.GetRoleById(l.ctx, &core.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	data := l.svcCtx.Casbin.GetFilteredPolicy(0, *roleData.Code)
	resp = &types.ApiAuthorityListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = uint64(len(data))
	for _, v := range data {
		resp.Data.Data = append(resp.Data.Data, types.ApiAuthorityInfo{
			Path:   v[1],
			Method: v[2],
		})
	}

	return resp, nil
}
