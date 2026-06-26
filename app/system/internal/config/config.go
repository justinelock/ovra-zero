package config

import (
	"ovra/toolkit/configshared"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	RestConf    rest.RestConf
	RpcConf     zrpc.RpcServerConf
	Tenant      configshared.TenantConfig
	Data        configshared.DataConfig
	JwtAuth     configshared.JwtAuthConfig
	ApiDecrypt  configshared.ApiDecryptConfig
	Captcha     configshared.CaptchaConfig
	Idempotency configshared.IdempotencyConfig
	Sign        configshared.SignConfig
	FileUpload  FileUploadConfig
}

// FileUploadConfig 本地文件上传根目录（投信海报等，对齐 Java file.upload.path）
type FileUploadConfig struct {
	Path string `json:",optional"`
}
