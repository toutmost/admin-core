package middleware

import (
	"context"
	"errors"
	"github.com/casbin/casbin/v2"
	"github.com/redis/go-redis/v9"
	"github.com/toutmost/admin-common/config"
	"github.com/toutmost/admin-common/enum/errorcode"
	"github.com/toutmost/admin-common/i18n"
	"github.com/toutmost/admin-common/utils/jwt"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"strings"
)

type AuthorityMiddleware struct {
	Cbn         *casbin.Enforcer
	Rds         redis.UniversalClient
	Trans       *i18n.Translator
	BanRoleData map[string]bool
}

func NewAuthorityMiddleware(cbn *casbin.Enforcer, rds redis.UniversalClient, trans *i18n.Translator, banRoleData map[string]bool) *AuthorityMiddleware {
	return &AuthorityMiddleware{
		Cbn:         cbn,
		Rds:         rds,
		Trans:       trans,
		BanRoleData: banRoleData,
	}
}

func (m *AuthorityMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the path
		obj := r.URL.Path
		// get the method
		act := r.Method
		// get the role id
		roleIds := strings.Split(r.Context().Value("roleId").(string), ",")

		// check the role status
		var isRoleNormal bool
		for _, v := range roleIds {
			if !m.BanRoleData[v] {
				isRoleNormal = true
			}
		}

		if !isRoleNormal {
			httpx.Error(w, errorx.NewApiErrorWithoutMsg(http.StatusUnauthorized))
			return
		}

		// check jwt blacklist
		jwtResult, err := m.Rds.Get(context.Background(), config.RedisTokenPrefix+jwt.StripBearerPrefixFromToken(r.Header.Get("Authorization"))).Result()
		if err != nil && !errors.Is(err, redis.Nil) {
			logx.Errorw("redis error in jwt", logx.Field("detail", err.Error()))
			httpx.Error(w, errorx.NewApiError(http.StatusInternalServerError, err.Error()))
			return
		}

		if jwtResult == "1" {
			logx.Errorw("token in blacklist", logx.Field("detail", r.Header.Get("Authorization")))
			httpx.Error(w, errorx.NewApiErrorWithoutMsg(http.StatusUnauthorized))
			return
		}

		result := batchCheck(m.Cbn, roleIds, act, obj)

		if result {
			logx.Infow("HTTP/HTTPS Request", logx.Field("UUID", r.Context().Value("userId").(string)),
				logx.Field("path", obj), logx.Field("method", act))
			next(w, r)
			return
		} else {
			logx.Errorw("the role is not permitted to access the API", logx.Field("roleId", roleIds),
				logx.Field("path", obj), logx.Field("method", act))
			httpx.Error(w, errorx.NewCodeError(errorcode.PermissionDenied, m.Trans.Trans(
				context.WithValue(context.Background(), "lang", r.Header.Get("Accept-Language")),
				i18n.PermissionDeny)))
			return
		}
	}
}

func batchCheck(cbn *casbin.Enforcer, roleIds []string, act, obj string) bool {
	var checkReq [][]any
	for _, v := range roleIds {
		checkReq = append(checkReq, []any{v, obj, act})
	}

	result, err := cbn.BatchEnforce(checkReq)
	if err != nil {
		logx.Errorw("Casbin enforce error", logx.Field("detail", err.Error()))
		return false
	}

	for _, v := range result {
		if v {
			return true
		}
	}

	return false
}
