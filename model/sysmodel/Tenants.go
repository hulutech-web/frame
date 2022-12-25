package sysmdel

import (
	"github.com/hulutech-web/frame/helpers/m"
	"github.com/hulutech-web/frame/helpers/zone"
	"github.com/hulutech-web/frame/model"
)

type Tenants struct {
	ID                 int64      `gorm:"column:id;primary_key;auto_increment"`
	UserId             int64      `gorm:"column:user_id;type:int unsigned;not null"`
	TenantsId          int64      `gorm:"column:tenants_id;type:int unsigned;not null"`
	Host               string     `gorm:"column:host;type:varchar(255);not null"`
	DriverName         string     `gorm:"column:drivername;type:varchar(255);not null"`
	Port               int64      `gorm:"column:port;type:int unsigned;not null"`
	Prefix             string     `gorm:"column:prefix;type:varchar(255);not null"`
	Dbname             string     `gorm:"column:dbname;type:varchar(255);not null"`
	Dbuser             string     `gorm:"column:dbuser;type:varchar(255);not null"`
	Charset            string     `gorm:"column:charset;type:varchar(255);not null"`
	Collation          string     `gorm:"column:collation;type:varchar(255);not null"`
	SetmaxIdleconns    int64      `gorm:"column:setmaxIdleconns;type:int unsigned;not null"`
	Setmaxopenconns    int64      `gorm:"column:setmaxopenconns;type:int unsigned;not null"`
	Setconnmaxlifetime int64      `gorm:"column:setconnmaxlifetime;type:int unsigned;not null"`
	Password           string     `gorm:"column:password;type:varchar(255);not null"`
	CreatedAt          zone.Time  `gorm:"column:created_at"`
	UpdatedAt          zone.Time  `gorm:"column:updated_at"`
	DeletedAt          *zone.Time `gorm:"column:deleted_at"`
	model.BaseModel
}

func (tenant *Tenants) TableName() string {
	return tenant.SetTableName("tenants_db")
}

func (tenant *Tenants) SetPortAttribute(value interface{}) {
	if tenant.Port == 0 {
		tenant.Port = 3306
	}
}

func (tenant *Tenants) ObjArr(filterArr []model.Filter, sortArr []model.Sort, limit int, withTrashed bool) (interface{}, error) {
	var outArr []Tenants
	if err := m.H().Q(filterArr, sortArr, limit, withTrashed).Find(&outArr).Error; err != nil {
		return nil, err
	}
	return outArr, nil
}

func (tenant *Tenants) ObjArrPaginate(c model.Context, perPage uint, filterArr []model.Filter, sortArr []model.Sort, limit int, withTrashed bool) (pagination model.Pagination, err error) {
	var outArr []Tenants
	filter := model.Model(*m.H().Q(filterArr, sortArr, limit, withTrashed))
	return filter.Paginate(&outArr, c, int64(perPage))
}
