package component

import (
	"powergate.com/hr_timesheet/cmd/server/configs"
)

type AppContext interface {
	SecretKey() string
	GetAppConfig() configs.Config
}

type appCtx struct {
	secretKey string
	configs   configs.Config
}

func NewAppContext(configs configs.Config) *appCtx {
	return &appCtx{configs: configs}
}

func (ctx *appCtx) GetAppConfig() configs.Config {
	return ctx.configs
}
func (ctx *appCtx) SecretKey() string {
	return ctx.secretKey
}
