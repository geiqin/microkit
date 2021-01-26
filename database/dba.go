package database

import (
	"context"
	"github.com/geiqin/microkit/auth"
	"gorm.io/gorm"
	"sync"
)

var mu sync.Mutex

//连接数据库（包括公共库和私有库）
func ConnectDB(ctx context.Context, appFlag string, appPrivate bool) *gorm.DB {
	mu.Lock()
	defer mu.Unlock()
	cfg := *GetConnectCfg(appFlag)

	if appPrivate {
		storeId := auth.GetStoreId(ctx)
		if storeId > 0 {
			cfg.Database = auth.GetStoreFlag(storeId)
		}
	}
	return DbPools(&cfg)
}
