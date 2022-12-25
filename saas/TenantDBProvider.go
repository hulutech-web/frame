package saas

import (
	"database/sql"
	"fmt"
	sysmdel "gitee.com/hulutech/frame/model/sysmodel"
	"sync"

	"gitee.com/hulutech/frame/request"
	"gorm.io/gorm"
	// "gitee.com/hulutech/frame/helpers/cache"
	// "gitee.com/hulutech/frame/helpers/debug"
)

type TenantInfoDB struct {
	DbName     string   // 数据库连接名称
	Tid        int64    // 租户id
	UserId     int64    // 用户id
	ConnStr    string   // 连接串
	DriverName string   // 驱动名称
	Db         *gorm.DB // 数据库连接对象
	Dber       *sql.DB  // 数据库连接对象
	Prefix     string
}

var (
	TenantsDBMutex  sync.Mutex
	TenantsDBMapsV2 = make(map[int64]*TenantInfoDB)
)

// SetDb 上下文连接数据库
func SetDb(c request.Context) (tenant *gorm.DB) {

	TenantsId, _ := c.Get("TenantsId")

	UserId, _ := c.UserId()

	tenantDB, _ := c.Get("TenantsDB")

	TenantsDB := tenantDB.(sysmdel.Tenants)

	itemDb := TenantInfoDB{}

	itemDb.DbName = TenantsDB.Dbname

	itemDb.Tid = TenantsDB.TenantsId

	itemDb.Prefix = TenantsDB.Prefix

	// 判断 是否同一租户
	if TenantsId != TenantsDB.TenantsId {
		panic("TenantsId == TenantsDB.TenantsId")
	}

	if _, ok := TenantsDBMapsV2[itemDb.Tid]; ok {
		ThemTenantsDB := TenantsDBMapsV2[itemDb.Tid]
		if ThemTenantsDB.UserId == UserId && ThemTenantsDB.Tid == TenantsDB.TenantsId && ThemTenantsDB.DriverName == TenantsDB.DriverName && ThemTenantsDB.DbName == TenantsDB.Dbname {
			return TenantsDBMapsV2[itemDb.Tid].Db
		}
	}

	var mysqlStr string
	var dber *sql.DB
	tenant, dber, mysqlStr = MysqlOpenDb(TenantsDB)

	itemDb.DriverName = TenantsDB.DriverName
	itemDb.Db = tenant
	itemDb.Dber = dber
	itemDb.ConnStr = mysqlStr

	TenantsDBMutex.Lock()

	defer TenantsDBMutex.Unlock()

	TenantsDBMapsV2[itemDb.Tid] = &itemDb

	return tenant

}

// 通过租户id 连接数据库
// SetTenantDB 上下文连接数据库
func SetTenantDB(TenantId int64) (db *gorm.DB) {

	var TenantsDB sysmdel.Tenants

	tenantDbWhere := &sysmdel.Tenants{
		ID: TenantId,
	}

	TenantsDB.DB().Where(tenantDbWhere).Find(&TenantsDB)

	itemDb := TenantInfoDB{}
	itemDb.DbName = TenantsDB.Dbname
	itemDb.Tid = TenantsDB.TenantsId
	itemDb.UserId = TenantsDB.UserId
	itemDb.Prefix = TenantsDB.Prefix
	// 判断 是否同一租户
	if TenantId != TenantsDB.TenantsId {
		panic("Illegal request!")
	}
	if _, ok := TenantsDBMapsV2[itemDb.Tid]; ok {
		ThemTenantsDB := TenantsDBMapsV2[itemDb.Tid]
		if ThemTenantsDB.UserId == TenantsDB.UserId && ThemTenantsDB.Tid == TenantsDB.TenantsId && ThemTenantsDB.DriverName == TenantsDB.DriverName && ThemTenantsDB.DbName == TenantsDB.Dbname {
			return TenantsDBMapsV2[itemDb.Tid].Db
		}
	}

	var mysqlStr string
	var dber *sql.DB
	db, dber, mysqlStr = MysqlOpenDb(TenantsDB)

	itemDb.DriverName = TenantsDB.DriverName
	itemDb.Db = db
	itemDb.Dber = dber
	itemDb.ConnStr = mysqlStr

	TenantsDBMutex.Lock()
	defer TenantsDBMutex.Unlock()

	TenantsDBMapsV2[itemDb.Tid] = &itemDb
	return db

}

//连接DB
func MysqlOpenDb(TenantsDB sysmdel.Tenants) (tenant *gorm.DB, dber *sql.DB, mysqlStr string) {
	var db *gorm.DB
	mysqlStr = fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true&charset=%s&collation=%s",
		TenantsDB.Dbuser,
		TenantsDB.Password,
		TenantsDB.Host,
		TenantsDB.Port,
		TenantsDB.Dbname,
		TenantsDB.Charset,
		TenantsDB.Collation,
	)
	db, dber = OpenDB(mysqlStr, int(TenantsDB.SetmaxIdleconns), int(TenantsDB.Setmaxopenconns), TenantsDB.Prefix)
	return db, dber, mysqlStr
}

func Close(dber *sql.DB) {
	dber.Close()
}
