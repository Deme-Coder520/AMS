package routers

import (
	"AMS/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/login", &controllers.UserController{},"get:ShowLogin;post:HandleLogin")
    beego.Router("/register", &controllers.UserController{},"get:ShowRegister;post:HandleRegister")
    beego.Router("/index", &controllers.ArticleController{},"get,post:ShowIndex")
    beego.Router("/addArticle", &controllers.ArticleController{},"get:ShowAdd;post:HandleAdd")
    beego.Router("/content", &controllers.ArticleController{},"get:ShowContent")
    beego.Router("/edit", &controllers.ArticleController{},"get:ShowEdit;post:Edit")
}
