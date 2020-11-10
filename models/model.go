package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

/*数据库表结构的设计应该放在model文件夹下面*/

// 创建一个model结构体，方便创表
type UserInfo struct {
	Id int
	Name string
	Password string
}

// 连接数据库
func init() {
	// 1.设置数据库基本信息，参数二：驱动的名称"mysql""sqlite""postgres"
	err := orm.RegisterDataBase("default","mysql","root:root@tcp(127.0.0.1:3306)/db?charset=utf8",30)
	if err != nil {
		beego.Info("连接数据库出错啦！")
		return
	}
	beego.Info("成功连接数据库~")
	// 2.映射model 数据
	orm.RegisterModel(new(UserInfo))
	// 3.生成表,参数一:数据库别名；参数二：是否强制更新表结构（若表结构该表需要切换成true）；参数三：创建过程在终端是否可见
	err = orm.RunSyncdb("default",false,true)
	if  err != nil {
		beego.Info("生成表出错了...")
		return
	}
	beego.Info("创建表成功~")
}
