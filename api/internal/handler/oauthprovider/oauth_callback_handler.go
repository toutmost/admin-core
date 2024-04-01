package oauthprovider

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/toutmost/admin-core/api/internal/logic/oauthprovider"
	"github.com/toutmost/admin-core/api/internal/svc"
)

// swagger:route get /oauth/login/callback oauthprovider OauthCallback
//
// Oauth log in callback route | Oauth 登录返回调接口
//
// Oauth log in callback route | Oauth 登录返回调接口
//
// Responses:
//  200: CallbackResp

func OauthCallbackHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := oauthprovider.NewOauthCallbackLogic(r.Context(), svcCtx)
		resp, err := l.OauthCallback()
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
