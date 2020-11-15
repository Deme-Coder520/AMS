package controllers

import (
	"AMS/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UserController struct {
	 beego.Controller
}

/*get和post的区别：
get和post本质上都是tcp链接，但由于http的规定和浏览器的限制，导致它们在应用过程中体现出了一些不同
1.GET在浏览器回退时对数据是无影响的，而POST会再次提交请求。
2.Get是不安全的，因为在传输过程，数据被放在请求的URL中；Post的所有操作对用户来说都是不可见的。
3.get的传输的数据较小，受URL长度的限制，而post传输的数据较大，一般默认为不受限制
4.GET请求参数会被完整保留在浏览器历史记录里，而POST中的参数不会被保留。
5.对参数的数据类型，GET只接受ASCII字符，而POST没有限制。*/

// ShowRegister 显示注册页面
func (u *UserController) ShowRegister() {
	u.TplName = "register.html"
}

// HandleRegister 处理注册业务
func (u *UserController) HandleRegister() {
	// 注册功能实现
	// 1.拿到前段传过来的用户数据
	userName := u.GetString("username")
	pwd := u.GetString("password")
	// 2.对用户数据进行校验
	if userName == "" || pwd == "" {
		beego.Info("用户名或密码不能为空！")
		//r.Redirect("/register",302)//请求重定向
		u.TplName = "register.html"
		return
	}
	// 3.校验通过，将数据插入到数据库
	o := orm.NewOrm()
	user := models.UserInfo{Name:userName,Password:pwd}
	_,err := o.Insert(&user)
	if err != nil {
		beego.Info("用户名或密码不能为空！")
		//r.Redirect("/register",302)//请求重定向
		u.TplName = "register.html"
		return
	}
	// 4.跳转到登录界面,两种方式：Redirect()速度快但是不能传输数据，TplName可以传输数据
	//u.TplName = "login.html"
	u.Redirect("/login",302)
}

// ShowLogin 展示登录界面
func (u *UserController) ShowLogin(){
	u.TplName = "login.html"
}

// HandleLogin 处理登录业务
func (u *UserController) HandleLogin() {
	// 实现登录功能
	// 1.获取前端传过来的数据
	userName := u.GetString("username")
	pwd := u.GetString("password")
	// 2.判断数据是否合法
	if userName == "" || pwd == "" {
		beego.Info("用户名或密码不能为空！")
		//l.Redirect("/login",302)
		u.TplName = "login.html"
		return
	}
	// 3.查询用户是否存在数据库中
	o := orm.NewOrm()
	user := models.UserInfo{Name:userName}
	err := o.Read(&user,"Name")
	if err != nil {
		beego.Info("用户不存在，请先注册！")
		//l.Redirect("/login",302)
		u.TplName = "login.html"
		return
	}
	if pwd != user.Password || userName != user.Name {
		// 判断用户密码是否正确
		beego.Info("用户名或密码错误！！")
		//l.Redirect("/login",302)
		u.TplName = "login.html"
		return
	}
	// 4.跳转指定界面
	u.Redirect("/index",302)
	//u.TplName = "index.html"
}

