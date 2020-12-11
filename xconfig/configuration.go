package xconfig

import (
	"github.com/geiqin/gotools/database"
	"github.com/geiqin/microkit/auth"
	"github.com/geiqin/microkit/cache"
	"github.com/geiqin/microkit/session"
	"os"
)

var conf *Configuration

//注册到配置中心
func Register(names ...string) {
	defNames := []string{"app", "database"}
	if names != nil {
		defNames = append(defNames, names...)
	}

	address := os.Getenv("SAAS_CONFIG_ADDRESS")
	if address == "" {
		panic("unset SAAS_CONFIG_ADDRESS env")
	}

	mgr := NewConfigManager(address, defNames)
	conf = mgr.Load()
}

//获取全部配置
func GetConfig() *Configuration {
	return conf
}

//获取数据库配置
func GetDatabaseCfg(name string) *database.DbConfig {
	cfg := conf.DatabaseList[name]
	return cfg
}

//获取文件储存配置
func GetFilesystemCfg(name string) *FileSystemInfo {
	cfg := conf.FileSystemList[name]
	return cfg
}

//获取店铺数据库配置
func GetStoreDatabaseCfg(name string, storeFlag string) *database.DbConfig {
	cfg := *conf.DatabaseList[name]
	if &cfg != nil {
		cfg.Database = storeFlag
	}
	return &cfg
}

//获取Redis配置
func GetRedisCfg(name string) *cache.RedisConfig {
	cfg := conf.RedisList[name]
	return cfg
}

//获取缓存配置
func GetCacheCfg() *cache.RedisConfig {
	cfg := conf.RedisList["cache"]
	return cfg
}

//获取会话配置
func GetSessionCfg() *session.SessConfig {
	if conf.SessionInfo != nil {
		conf.SessionInfo.Provider = GetRedisCfg("session")
	}
	return conf.SessionInfo
}

//获取令牌配置
func GetTokenCfg(name string) *auth.TokenConfig {
	cfg := conf.TokenList[name]
	return cfg
}

//获取微信支付配置
func GetWxPayCfg() *WxPayInfo {
	cfg := conf.WxPayInfo
	return cfg
}

//获取支付宝支付配置
func GetAliPayCfg() *AliPayInfo {
	cfg := conf.AliPayInfo
	return cfg
}
