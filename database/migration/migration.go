package migration

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/hulutech-web/frame/database"

	"github.com/hulutech-web/frame/config"
)

type Migration struct {
	ID        int64  `gorm:"column:id;primary_key;auto_increment;"`
	Migration string `gorm:"column:migration;type:varchar(255)"`
	Batch     int64  `gorm:"column:batch;"`
}

// 建立migration表
func (m *Migration) up(db *gorm.DB) {
	tx := db.Begin()
	{
		tx.Migrator().CreateTable(&m)
	}
	tx.Commit()
}

func (m *Migration) TableName() string {
	tableName := config.GetString("database.migrations")
	if tableName == "" {
		panic("migrations table name parse error")
	}

	tableNameWithPrefix := fmt.Sprintf("%s%s", database.Prefix(), tableName)
	return tableNameWithPrefix
}

func (m *Migration) Name() string {
	return m.Migration
}
