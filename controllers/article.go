package controllers

import (
	"github.com/astaxie/beego"
)

type ArticleController struct {
	beego.Controller
}

// ShowIndex 展示首页
func (a *ArticleController) ShowIndex(){
	a.TplName = "index.html"
}