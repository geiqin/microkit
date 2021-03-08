package database

import (
	"context"
	"github.com/geiqin/microkit/auth"
	"gorm.io/gorm"
	"sync"
)

var mu sync.Mutex

//连接数据库（包括公共库和私有库）
func ConnectDB(ctx context.Context, appFlag string, appPrivate bool, storeDbPrefix ...string) *gorm.DB {
	mu.Lock()
	defer mu.Unlock()
	cfg := *GetConnectCfg(appFlag)

	if appPrivate {
		p := "go_store_"
		if storeDbPrefix != nil {
			p = storeDbPrefix[0]
		}
		storeId := auth.GetStoreId(ctx)
		if storeId > 0 {
			cfg.Database = auth.GetStoreFlag(storeId, p)
		}
	}
	return DbPools(&cfg)
}
