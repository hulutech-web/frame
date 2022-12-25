package saas

import (
	"database/sql"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	dbconfig "gitee.com/hulutech/frame/config"
)

func OpenDB(dsn string, maxIdleConns, maxOpenConns int, Prefix string) (db *gorm.DB, dber *sql.DB) {
	config := &gorm.Config{}
	var err error
	var sqldb *sql.DB
	if Prefix == "" {
		Prefix = dbconfig.GetString("DB_PREFIX")
	}
	if config.NamingStrategy == nil {
		config.NamingStrategy = schema.NamingStrategy{
			TablePrefix:   Prefix,
			SingularTable: true,
		}
	}
	// 连接DB
	if db, err = gorm.Open(mysql.Open(dsn), config); err != nil {
		log.Errorf("opens database failed: %s", err.Error())
		return
	}
	// 连接池
	if sqldb, err = db.DB(); err == nil {
		sqldb.SetMaxIdleConns(maxIdleConns)
		sqldb.SetMaxOpenConns(maxOpenConns)
	} else {
		log.Error(err)
	}

	return db, sqldb
}

func DbClose(dber *sql.DB) {
	dber.Close()
}
