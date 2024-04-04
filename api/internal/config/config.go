package config

import (
	"github.com/toutmost/admin-common/config"
	"github.com/toutmost/admin-common/i18n"
	"github.com/toutmost/admin-common/plugins/casbin"
	"github.com/toutmost/admin-common/utils/captcha"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth         rest.AuthConf
	RedisConf    config.RedisConf
	CoreRpc      zrpc.RpcClientConf
	McmsRpc      zrpc.RpcClientConf
	Captcha      captcha.Conf
	DatabaseConf config.DatabaseConf
	CasbinConf   casbin.CasbinConf
	I18nConf     i18n.Conf
	ProjectConf  ProjectConf
	CROSConf     config.CROSConf
}

type ProjectConf struct {
	DefaultRoleId           uint64 `json:",default=1"`
	DefaultDepartmentId     uint64 `json:",default=1"`
	DefaultPositionId       uint64 `json:",default=1"`
	EmailCaptchaExpiredTime int    `json:",default=600"`
	SmsTemplateId           string `json:",optional"`
	SmsAppId                string `json:",optional"`
	SmsSignName             string `json:",optional"`
	SmsParamsType           string `json:",default=json,options=[json,array]"`
	RegisterVerify          string `json:",default=captcha,options=[disable,captcha,email,sms,sms_or_email]"`
	LoginVerify             string `json:",default=captcha,options=[captcha,email,sms,sms_or_email,all]"`
	ResetVerify             string `json:",default=email,options=[email,sms,sms_or_email]"`
	AllowInit               bool   `json:",default=true"`
}
