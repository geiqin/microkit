package app

import (
	"github.com/geiqin/gotools/helper"
	//"github.com/geiqin/supports/auth"
	"github.com/geiqin/microkit/cache"
	"github.com/geiqin/microkit/session"
	"log"
	"os"
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

	session.Load(&session.SessConfig{
		Driver:      "redis",
		CookieName:  os.Getenv("QIN_SESSION_COOKIE_NAME"),
		MaxLifeTime: 3600,
		Provider:    &session.RedisProviderConfig{
			Host:     os.Getenv("QIN_REDIS_HOST"),
			Port:    helper.StringToInt(os.Getenv("QIN_REDIS_PORT")),
			Username: os.Getenv("QIN_REDIS_USERNAME"),
			Password: os.Getenv("QIN_REDIS_PASSWORD"),
			Database: 0,
		},
	})

	cache.Load(&cache.RedisConfig{
		Host:     os.Getenv("QIN_REDIS_HOST"),
		Port:    helper.StringToInt(os.Getenv("QIN_REDIS_PORT")),
		Username: os.Getenv("QIN_REDIS_USERNAME"),
		Password: os.Getenv("QIN_REDIS_PASSWORD"),
		Database: 1,
	})

	//database.Load(opt.Flag)
	//auth.Load()
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
