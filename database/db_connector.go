package database

import (
	"database/sql"
	"time"

	"github.com/hulutech-web/frame/config"
	"github.com/hulutech-web/frame/database/driver"
	"github.com/hulutech-web/frame/helpers/zone"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)


var db *gorm.DB
var dber databaser



func Initialize() {
	dber, db = setv2Connection("default")
}


func setv2Connection(conn string) (dber databaser, sqlDb *gorm.DB) {

	if conn == "" {
		panic("database connection parse error")
	}
	conn = config.GetString("database." + conn)

	switch conn {
	//mysql 驱动
	case "mysql":
		dber = driver.NewMysql(conn)
		Db, err := sql.Open("mysql", dber.ConnectionArgs())
		sqlDb, err = gorm.Open(mysql.New(mysql.Config{
			DSN:                       dber.ConnectionArgs(),
			DefaultStringSize:         256,   // string 类型字段的默认长度
			DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
			DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
			DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
			SkipInitializeWithVersion: false, // 根据版本自动配置
			Conn:                      Db,
		}), &gorm.Config{})
		err = Db.Ping()
		if err != nil {
			panic("failed to connect database by ping")
		}
		Db.SetConnMaxLifetime(time.Hour)
		Db.SetMaxIdleConns(config.GetInt("database.max_idle_connections"))
		Db.SetMaxOpenConns(config.GetInt("database.max_open_connections"))
		Db.SetConnMaxLifetime(zone.Duration(config.GetInt("database.max_life_seconds")) * zone.Second)

		return dber, sqlDb
		break
	//mssql 驱动
	case "mssql":
		dber = driver.NewMssql(conn)
		sqlDb, err := gorm.Open(sqlserver.Open(dber.ConnectionArgs()), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		db, err := sqlDb.DB()
		err = db.Ping()
		if err != nil {
			panic("failed to connect database by ping")
		}
		db.SetConnMaxLifetime(time.Hour)
		db.SetMaxIdleConns(config.GetInt("database.max_idle_connections"))
		db.SetMaxOpenConns(config.GetInt("database.max_open_connections"))
		db.SetConnMaxLifetime(zone.Duration(config.GetInt("database.max_life_seconds")) * zone.Second)
		return dber, sqlDb
		break
	default:
		panic("incorrect database connection provided")
	}
	return
}

func Connection(conn string) (db *gorm.DB) {
	_, db = setv2Connection(conn)
	return db
}

func DB() *gorm.DB {
	return db
}




func Prefix() string {
	return dber.Prefix()
}