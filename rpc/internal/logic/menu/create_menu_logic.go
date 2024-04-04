package menu

import (
	"context"
	"github.com/toutmost/admin-common/enum/common"
	"github.com/toutmost/admin-common/i18n"
	"github.com/toutmost/admin-core/rpc/ent/menu"
	"github.com/toutmost/admin-core/rpc/internal/utils/dberrorhandler"

	"github.com/toutmost/admin-core/rpc/internal/svc"
	"github.com/toutmost/admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMenuLogic {
	return &CreateMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateMenuLogic) CreateMenu(in *core.MenuInfo) (*core.BaseIDResp, error) {
	// get parent level
	var menuLevel uint32
	if *in.ParentId != common.DefaultParentId {
		m, err := l.svcCtx.DB.Menu.Query().Where(menu.IDEQ(*in.ParentId)).First(l.ctx)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
		}

		menuLevel = m.MenuLevel + 1
	} else {
		menuLevel = 1
	}

	result, err := l.svcCtx.DB.Menu.Create().
		SetNotNilMenuLevel(&menuLevel).
		SetNotNilMenuType(in.MenuType).
		SetNotNilParentID(in.ParentId).
		SetNotNilPath(in.Path).
		SetNotNilName(in.Name).
		SetNotNilRedirect(in.Redirect).
		SetNotNilComponent(in.Component).
		SetNotNilSort(in.Sort).
		SetNotNilDisabled(in.Disabled).
		SetNotNilServiceName(in.ServiceName).
		// meta
		SetNotNilTitle(in.Meta.Title).
		SetNotNilIcon(in.Meta.Icon).
		SetNotNilHideMenu(in.Meta.HideMenu).
		SetNotNilHideBreadcrumb(in.Meta.HideBreadcrumb).
		SetNotNilIgnoreKeepAlive(in.Meta.IgnoreKeepAlive).
		SetNotNilHideTab(in.Meta.HideTab).
		SetNotNilFrameSrc(in.Meta.FrameSrc).
		SetNotNilCarryParam(in.Meta.CarryParam).
		SetNotNilHideChildrenInMenu(in.Meta.HideChildrenInMenu).
		SetNotNilAffix(in.Meta.Affix).
		SetNotNilDynamicLevel(in.Meta.DynamicLevel).
		SetNotNilRealPath(in.Meta.RealPath).
		Save(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess}, nil
}
