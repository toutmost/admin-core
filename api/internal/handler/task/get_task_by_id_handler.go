package task

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/toutmost/admin-core/api/internal/logic/task"
	"github.com/toutmost/admin-core/api/internal/svc"
	"github.com/toutmost/admin-core/api/internal/types"
)

// swagger:route post /task task GetTaskById
//
// Get task by ID | 通过ID获取定时任务
//
// Get task by ID | 通过ID获取定时任务
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDReq
//
// Responses:
//  200: TaskInfoResp

func GetTaskByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := task.NewGetTaskByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetTaskById(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
