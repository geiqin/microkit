package xconfig

import (
	"github.com/geiqin/gotools/database"
	"github.com/geiqin/microkit/auth"
	"github.com/geiqin/microkit/cache"
	"github.com/geiqin/microkit/session"
	"github.com/micro/go-micro/v2/config"
	grpcConfig "github.com/micro/go-plugins/config/source/grpc/v2"
	"log"
)

type ConfigManger struct {
	address string
	conf    *Configuration
	names   []string
}

func (b *ConfigManger) readGrpc(name string) (config.Config, error) {
	log.Println("config addr:", b.address)
	source := grpcConfig.NewSource(
		grpcConfig.WithAddress(b.address),
		grpcConfig.WithPath("/"+name),
	)

	// create new config
	conf, _ := config.NewConfig()

	// load the source into config
	if err := conf.Load(source); err != nil {
		return nil, err
	}

	return conf, nil

}

func (b *ConfigManger) Load() *Configuration {
	for _, name := range b.names {
		conf, err := b.readGrpc(name)
		if err != nil {
			log.Fatal("read_cfg_error::", err)
		}

		switch name {
		case "app":
			sess := &session.SessConfig{}
			sessCfg := conf.Get("session")
			sessCfg.Scan(sess)
			b.conf.SessionInfo = sess
			break
		case "database":
			dbA := conf.Get("connections")
			dbB := conf.Get("redis")
			var dbAList map[string]*database.DbConfig
			var dbBList map[string]*cache.RedisConfig
			dbA.Scan(dbAList)
			dbB.Scan(dbBList)
			b.conf.DatabaseList = dbAList
			b.conf.RedisList = dbBList
			break
		case "filesystem":
			var cloudList map[string]*FileSystemInfo
			cloudCfg := conf.Get("connections")
			cloudCfg.Scan(cloudList)
			b.conf.FileSystemList = cloudList
			break
		case "token":
			var tkList map[string]*auth.TokenConfig
			cloudCfg := conf.Get("tokens")
			cloudCfg.Scan(tkList)
			b.conf.TokenList = tkList
			break
		case "payment":
			var wx *WxPayInfo
			wxCfg := conf.Get("providers", "weixin")
			wxCfg.Scan(wx)
			b.conf.WxPayInfo = wx
			break
		}
	}
	return b.conf
}

func NewConfigManager(address string, names []string) *ConfigManger {
	obj := &ConfigManger{
		address: address,
		conf:    &Configuration{},
		names:   names,
	}
	return obj
}
