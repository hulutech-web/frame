package sysmdel

import (
	"gorm.io/gorm"
	"time"
)

/*
*
该模型用来通过jwt中间件动态生成token,以及添加拥有者模型名称，模型id，以及登录模型的token(jwt生成的)
*/
type PersonalAccessToken struct {
	gorm.Model
	OwnerType string `gorm:"column:owner_type;type:varchar(255);not null;comment:'拥有者模型名称'"`
	OwnerId   uint  `gorm:"column:owner_id;type:bigint;not null;comment:'拥有者模型id'"`
	Token     string `gorm:"column:token;type:varchar(255);not null;comment:'登录模型的token(jwt生成的)'"`
	//默认值为当前时间
	LastUsedAt time.Time `gorm:"column:last_used_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:'最后使用时间'"`
}

