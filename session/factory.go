package session

import (
	"context"
	"fmt"
	"github.com/geiqin/microkit/auth"
	"github.com/geiqin/xconfig/model"
	"log"
)

var globalSessionManager *Manager

func Load(cnf *model.SessionInfo) {
	if cnf == nil {
		log.Println("load session config failed")
	}
	log.Println("load session config succeed")
	LoadRedis(cnf)
	newManager(cnf)
}

func newManager(cfg *model.SessionInfo) {
	var err error
	globalSessionManager, err = NewSessionManager(cfg.Driver, cfg.CookieName, cfg.MaxLifeTime)
	if err != nil {
		fmt.Println(err)
		return
	}
	if cfg.Driver == "memory" {
		go globalSessionManager.GC()
	}
}

func GetSession(ctx context.Context) (session Session) {
	session = globalSessionManager.SessionStart(GetSessionId(ctx))
	return session
}

func GetSessionById(sessionId string) (session Session) {
	session = globalSessionManager.SessionStart(sessionId)
	return session
}

func Destroy(ctx context.Context) {
	globalSessionManager.SessionDestroy(GetSessionId(ctx))
}

func GetSessionId(ctx context.Context) string {
	return auth.GetSessionId(ctx)
}
