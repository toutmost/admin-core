package department

import (
	"context"
	"github.com/toutmost/admin-common/utils/pointy"
	"github.com/toutmost/admin-core/rpc/ent/department"
	"github.com/toutmost/admin-core/rpc/ent/predicate"
	"github.com/toutmost/admin-core/rpc/internal/utils/dberrorhandler"

	"github.com/toutmost/admin-core/rpc/internal/svc"
	"github.com/toutmost/admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDepartmentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDepartmentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDepartmentListLogic {
	return &GetDepartmentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDepartmentListLogic) GetDepartmentList(in *core.DepartmentListReq) (*core.DepartmentListResp, error) {
	var predicates []predicate.Department
	if in.Name != nil {
		predicates = append(predicates, department.NameContains(*in.Name))
	}
	if in.Leader != nil {
		predicates = append(predicates, department.LeaderContains(*in.Leader))
	}
	result, err := l.svcCtx.DB.Department.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &core.DepartmentListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &core.DepartmentInfo{
			Id:        &v.ID,
			CreatedAt: pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt: pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			Status:    pointy.GetPointer(uint32(v.Status)),
			Sort:      &v.Sort,
			Name:      &v.Name,
			Ancestors: &v.Ancestors,
			Leader:    &v.Leader,
			Phone:     &v.Phone,
			Email:     &v.Email,
			Remark:    &v.Remark,
			ParentId:  &v.ParentID,
		})
	}

	return resp, nil
}
