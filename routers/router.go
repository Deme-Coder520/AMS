package routers

import (
	"AMS/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    // 登录
    beego.Router("/login", &controllers.UserController{},"get:ShowLogin;post:HandleLogin")
    // 注册
    beego.Router("/register", &controllers.UserController{},"get:ShowRegister;post:HandleRegister")
    // 展示首页
    beego.Router("/index", &controllers.ArticleController{},"get,post:ShowIndex")
    // 添加文章
    beego.Router("/addArticle", &controllers.ArticleController{},"get:ShowAdd;post:HandleAdd")
    // 文章详情
    beego.Router("/content", &controllers.ArticleController{},"get:ShowContent")
    // 编辑文章
    beego.Router("/edit", &controllers.ArticleController{},"get:ShowEdit;post:Edit")
    // 删除文章
    beego.Router("/delete", &controllers.ArticleController{},"get:Delete")
	// 文章类型
	beego.Router("/addType",&controllers.ArticleController{},"get:ShowArtType;post:AddType")
}
