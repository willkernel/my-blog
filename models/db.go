package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"myblog/models/class"
)

// 使用orm连接数据库步骤
// 1.注册数据库驱动 RegisterDriver(name,type)
// 2.注册数据库 RegisterDataBase(aliasName,driveName,datasource,params...)
// 3.注册对象模型 RegisterModel(models...)
// 4.开启同步 RunSyncdb(name string,force bool,verbose bool)

func init() {
	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:APIDemo2016@tcp(localhost:3306)/myblog?charset=utf8")
	orm.RegisterModel(new(class.User))
	orm.RunSyncdb("default", false, true)
}
