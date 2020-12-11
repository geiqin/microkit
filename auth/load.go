package auth

import (
	"log"
)

var storeConf *TokenConfig
var userConf *TokenConfig

type TokenConfig struct {
	Issuer     string `json:"issuer"`
	Audience   string `json:"audience"`
	PrivateKey []byte `json:"private_key"`
	ExpireTime int    `json:"expire_time"`
}

func Load(storeConf *TokenConfig, userConf *TokenConfig) {
	if storeConf == nil {
		log.Println("load store_token config failed")
		return
	}
	log.Println("load store_token config succeed")

	if userConf == nil {
		log.Println("load user_token config failed")
		return
	}
	log.Println("load user_token config succeed")
}
