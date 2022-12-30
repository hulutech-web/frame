package sysmdel

import (
"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Username      string `gorm:"type:varchar(20);not null;unique;<-:create" json:"username"`
	Password      string `gorm:"type:varchar(100);column:password;not null;" json:"-"`
	Name          string `gorm:"type:varchar(20);not null" json:"name"`
	Avatar        string `gorm:"type:varchar(100)" json:"avatar"`
	RememberToken string `gorm:"type:varchar(100)" json:"-"`
}


