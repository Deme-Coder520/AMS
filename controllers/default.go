package controllers

import (
	"github.com/astaxie/beego"
)

// MainController 控制器
type MainController struct {
	beego.Controller
}

// Get 方法默认
func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
