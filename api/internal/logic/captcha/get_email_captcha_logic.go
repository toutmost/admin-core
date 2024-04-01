package captcha

import (
	"context"

	"github.com/toutmost/admin-core/api/internal/svc"
	"github.com/toutmost/admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEmailCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetEmailCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEmailCaptchaLogic {
	return &GetEmailCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetEmailCaptchaLogic) GetEmailCaptcha(req *types.EmailCaptchaReq) (resp *types.BaseMsgResp, err error) {
	// todo: add your logic here and delete this line

	return
}
