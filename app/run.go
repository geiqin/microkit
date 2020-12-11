package app

import (
	"github.com/geiqin/gotools/helper"
	"github.com/geiqin/microkit/auth"
	"github.com/geiqin/microkit/cache"
	"github.com/geiqin/microkit/session"
	"github.com/geiqin/microkit/xconfig"
	"log"
)

var appOption *Option

//var once sync.Once

type ConfigMode string

const (
	LocalMode ConfigMode = "local"
	CloudMode ConfigMode = "cloud"
)

type Option struct {
	Flag       string
	Private    bool
	ConfigMode ConfigMode
	ConfigPath string
}

func Run(flag string, private bool, option ...Option) {
	log.Println("app flag [" + flag + "] is running")

	opt := &Option{}
	if option != nil {
		opt = &option[0]
	}
	opt.Flag = flag
	opt.Private = private
	appOption = opt

	xSessionCnf := xconfig.GetSessionCfg()
	sessionCnf := &session.SessConfig{}
	helper.StructCopy(sessionCnf, xSessionCnf)
	session.Load(sessionCnf)

	xCacheCnf := xconfig.GetCacheCfg()
	cacheCnf := &cache.RedisConfig{}
	helper.StructCopy(cacheCnf, xCacheCnf)
	cache.Load(cacheCnf)

	xStoreConf := xconfig.GetTokenCfg("store")
	xUserConf := xconfig.GetTokenCfg("user")
	storeConf := &auth.TokenConfig{}
	userConf := &auth.TokenConfig{}

	helper.StructCopy(storeConf, xStoreConf)
	helper.StructCopy(userConf, xUserConf)

	auth.Load(storeConf, userConf)

}

func Flag() string {
	return appOption.Flag
}

func Private() bool {
	return appOption.Private
}

func GetOption() *Option {
	return appOption
}
