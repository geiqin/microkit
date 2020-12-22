package database

import (
	"fmt"
	"github.com/geiqin/xconfig/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

var dbConnections map[string]*model.DatabaseInfo

func Load(connections map[string]*model.DatabaseInfo) {
	dbConnections = connections
}

//获取数据库连接配置
func GetConnectCfg(name string, storeFlag ...string) *model.DatabaseInfo {
	cfg := dbConnections[name]
	if &cfg != nil && storeFlag != nil {
		cfg.Database = storeFlag[0]
	}
	return cfg
}

func CreateMysqlDB(cfg *model.DatabaseInfo) *gorm.DB {
	serverAddr := cfg.Host + ":" + cfg.Port

	//当前有效两种
	//connString := cfg.Username + ":" + cfg.Password + "@tcp(" + serverAddr + ")/" + cfg.Database + "?charset=utf8mb4&loc=Local"
	connString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&loc=Local&&parseTime=True", cfg.Username, cfg.Password, serverAddr, cfg.Database)

	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   cfg.Prefix, // table name prefix, table for `User` would be `t_users`
			SingularTable: false,      // use singular table name, table for `User` would be `user` with this option enabled
		},
	})

	if err != nil {
		log.Println("mysql database connection failed :", cfg.Database)
		return nil
	}

	return db
}
