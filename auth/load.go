package auth

import (
	"github.com/geiqin/xconfig/model"
	"log"
)

var tokenCfg *model.TokenInfo

func Load(tokenConf *model.TokenInfo) {
	tokenCfg = tokenConf
	if tokenCfg == nil {
		log.Println("load token config failed")
		return
	}
	log.Println("load token config succeed")
}
