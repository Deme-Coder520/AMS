package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

/*数据库表结构的设计应该放在model文件夹下面*/

// UserInfo 用户信息结构体
type UserInfo struct {
	Id int
	Name string `orm:"unique"`
	Password string
}

// Article 文章信息结构体
type Article struct {
	Id int `orm:"pk;auto"`
	ArtName string `orm:"size(48);unique"`
	ArtCreateAt time.Time `orm:"auto_now;type(datetime)"`
	ArtCount int `orm:"default(0)"`
	ArtContent string
	ArtImg string `orm:"null"`
}

// init 连接数据库
func init() {
	// 1.设置数据库基本信息，参数二：驱动的名称"mysql""sqlite""postgres"
	err := orm.RegisterDataBase("default","mysql","root:root@tcp(127.0.0.1:3306)/db?charset=utf8",30)
	if err != nil {
		beego.Info("connect sql fail")
		return
	}
	beego.Info("Connect sql success")
	// 2.映射model 数据
	orm.RegisterModel(new(UserInfo),new(Article))
	// 3.生成表,参数一:数据库别名；参数二：是否强制更新表结构（若表结构该表需要切换成true）；参数三：创建过程在终端是否可见
	err = orm.RunSyncdb("default",false,true)
	if  err != nil {
		beego.Info("Create table fail")
		return
	}
	beego.Info("Create table successful")
}
