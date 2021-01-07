package app

import (
	"github.com/geiqin/gotools/helper"
	"github.com/geiqin/microkit/auth"
	"github.com/geiqin/microkit/cache"
	"github.com/geiqin/microkit/database"
	"github.com/geiqin/microkit/session"
	"github.com/geiqin/xconfig/client"
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

	appCfg := client.GetAppConfig()
	log.Println("appCfg:", helper.JsonEncode(appCfg))
	log.Println("app_cfg:", helper.JsonEncode(appCfg.Token))
	databaseCfg := client.GetDatabaseConfig()
	connCfg := databaseCfg.Connections
	database.Load(connCfg)
	sessionCnf := appCfg.Session
	cacheCnf := databaseCfg.RedisList["cache"]
	sessionRedisCnf := databaseCfg.RedisList["session"]
	sessionCnf.Provider = sessionRedisCnf
	session.Load(sessionCnf)
	cache.Load(cacheCnf)

	auth.Load(appCfg.Token)

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
