package controllers

import (
	"AMS/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get(){
	l.TplName = "login.html"
}

func (l *LoginController) Post() {
	// 实现登录功能
	// 1.获取前端传过来的数据
	userName := l.GetString("userName")
	pwd := l.GetString("password")
	// 2.判断数据是否合法
	if userName == "" || pwd == "" {
		beego.Info("用户名或密码不能为空！")
		//l.Redirect("/login",302)
		l.TplName = "login.html"
		return
	}
	// 3.查询用户是否存在数据库中
	o := orm.NewOrm()
	user := models.UserInfo{Name:userName}
	err := o.Read(&user,"Name")
	if err != nil {
		beego.Info("用户不存在，请先注册！")
		//l.Redirect("/login",302)
		l.TplName = "login.html"
		return
	}
	if pwd != user.Password || userName != user.Name {
		// 判断用户密码是否正确
		beego.Info("用户名或密码错误！！")
		//l.Redirect("/login",302)
		l.TplName = "login.html"
		return
	}
	// 4.跳转指定界面
	//l.Ctx.WriteString("恭喜您，登录成功！")
	l.TplName = "index.html"
}