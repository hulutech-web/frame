package sysmdel

import (
	"gitee.com/hulutech/frame/helpers/zone"
	"gitee.com/hulutech/frame/model"
)

type UserToken struct {
	model.BaseModel
	ID        int64      `gorm:"column:id;primary_key;auto_increment"`
	Token     string     `gorm:"size:255;unique;not null" json:"token" form:"token"`
	UserId    int64      `gorm:"column:user_id;not null;index:idx_user_token_user_id;" json:"userId" form:"userId"`
	TenantsId int64      `gorm:"column:tenants_id;type:int unsigned;not null"`
	ExpiredAt int64      `gorm:"column:expired_at;not null" json:"expiredAt" form:"expiredAt"`
	Status    int        `gorm:"column:status;not null;index:idx_user_token_status" json:"status" form:"status"`
	CreatedAt *zone.Time `gorm:"not null" json:"created_at" form:"createTime"`
	UpdatedAt *zone.Time `gorm:"not null" json:"created_at" form:"createTime"`
}
