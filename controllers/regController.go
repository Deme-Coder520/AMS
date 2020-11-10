package controllers

import (
	"AMS/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type RegController struct {
	 beego.Controller
}

/*get和post的区别：
get和post本质上都是tcp链接，但由于http的规定和浏览器的限制，导致它们在应用过程中体现出了一些不同
1.GET在浏览器回退时对数据是无影响的，而POST会再次提交请求。
2.Get是不安全的，因为在传输过程，数据被放在请求的URL中；Post的所有操作对用户来说都是不可见的。
3.get的传输的数据较小，受URL长度的限制，而post传输的数据较大，一般默认为不受限制
4.GET请求参数会被完整保留在浏览器历史记录里，而POST中的参数不会被保留。
5.对参数的数据类型，GET只接受ASCII字符，而POST没有限制。*/

func (r *RegController) Get() {
	r.TplName = "register.html"
}

func (r *RegController) Post() {
	// 注册功能实现
	// 1.拿到前段传过来的用户数据
	userName := r.GetString("userName")
	pwd := r.GetString("password")
	// 2.对用户数据进行校验
	if userName == "" || pwd == "" {
		beego.Info("用户名或密码不能为空！")
		//r.Redirect("/register",302)//请求重定向
		r.TplName = "register.html"
		return
	}
	// 3.校验通过，将数据插入到数据库
	o := orm.NewOrm()
	user := models.UserInfo{Name:userName,Password:pwd}
	_,err := o.Insert(&user)
	if err != nil {
		beego.Info("用户名或密码不能为空！")
		//r.Redirect("/register",302)//请求重定向
		r.TplName = "register.html"
		return
	}
	// 4.跳转到登录界面,两种方式：Redirect()速度快但是不能传输数据，TplName可以传输数据
	r.TplName = "login.html"
}

