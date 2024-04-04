package menu

import (
	"context"
	"github.com/toutmost/admin-common/i18n"
	"github.com/toutmost/admin-common/utils/pointy"
	"github.com/toutmost/admin-core/rpc/types/core"

	"github.com/toutmost/admin-core/api/internal/svc"
	"github.com/toutmost/admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuListLogic {
	return &GetMenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetMenuListLogic) GetMenuList() (resp *types.MenuPlainInfoListResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetMenuList(l.ctx, &core.PageInfoReq{
		Page:     1,
		PageSize: 100,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.MenuPlainInfoListResp{}
	resp.Data.Total = data.Total
	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data, types.MenuPlainInfo{
			Id:                 v.Id,
			CreatedAt:          v.CreatedAt,
			UpdatedAt:          v.UpdatedAt,
			Trans:              pointy.GetPointer(l.svcCtx.Trans.Trans(l.ctx, *v.Meta.Title)),
			MenuType:           v.MenuType,
			Level:              v.Level,
			Path:               v.Path,
			Name:               v.Name,
			Redirect:           v.Redirect,
			Component:          v.Component,
			Sort:               v.Sort,
			ParentId:           v.ParentId,
			Title:              v.Meta.Title,
			Icon:               v.Meta.Icon,
			HideMenu:           v.Meta.HideMenu,
			HideBreadcrumb:     v.Meta.HideBreadcrumb,
			IgnoreKeepAlive:    v.Meta.IgnoreKeepAlive,
			HideTab:            v.Meta.HideTab,
			FrameSrc:           v.Meta.FrameSrc,
			CarryParam:         v.Meta.CarryParam,
			HideChildrenInMenu: v.Meta.HideChildrenInMenu,
			Affix:              v.Meta.Affix,
			DynamicLevel:       v.Meta.DynamicLevel,
			RealPath:           v.Meta.RealPath,
			Disabled:           v.Disabled,
			ServiceName:        v.ServiceName,
		})
	}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	return resp, nil
}
