package auth

import (
	"github.com/geiqin/xconfig/model"
	"log"
)

var storeConf *model.TokenInfo
var userConf *model.TokenInfo

/*
type TokenConfig struct {
	Issuer     string `json:"issuer"`
	Audience   string `json:"audience"`
	PrivateKey []byte `json:"private_key"`
	ExpireTime int    `json:"expire_time"`
}

*/

func Load(storeConf *model.TokenInfo, userConf *model.TokenInfo) {
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
