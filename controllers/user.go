package controllers

import (
	"AMS/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// UserController 自定义控制器
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
		u.Data["errMsg"] = "用户名或密码不能为空！"
		u.TplName = "register.html"
		return
	}
	// 3.校验通过，将数据插入到数据库
	o := orm.NewOrm()
	user := models.UserInfo{Name: userName, Password: pwd}
	err := o.Read(&user, "Name")
	if err == nil {
		u.Data["errMsg"] = "该用户名已被注册"
		u.TplName = "register.html"
		return
	}
	_, err = o.Insert(&user)
	if err != nil {
		beego.Info("添加数据失败")
		//r.Redirect("/register",302)//请求重定向
		u.TplName = "register.html"
		return
	}
	// 4.跳转到登录界面,两种方式：Redirect()速度快但是不能传输数据，TplName可以传输数据
	u.Redirect("/login", 302)
}

// ShowLogin 展示登录界面
func (u *UserController) ShowLogin() {
	u.TplName = "login.html"
	// 获取cookie，如果有，显示用户名，没有就显示空
	userName := u.Ctx.GetCookie("username")
	//fmt.Println(userName)
	//decName,_ := base64.StdEncoding.DecodeString(userName)
	if userName != "" {
		u.Data["username"] = userName
		u.Data["checkStatus"] = "checked"
	}
}

// HandleLogin 处理登录业务
func (u *UserController) HandleLogin() {
	// 实现登录功能
	// 1.获取前端传过来的数据
	userName := u.GetString("username")
	pwd := u.GetString("password")
	// 2.判断数据是否合法
	if userName == "" || pwd == "" {
		u.Data["errMsg"] = "用户名或密码不能为空！"
		//l.Redirect("/login",302)
		u.TplName = "login.html"
		return
	}
	// 3.查询用户是否存在数据库中
	o := orm.NewOrm()
	user := models.UserInfo{Name: userName}
	err := o.Read(&user, "Name")
	if err != nil {
		u.Data["errMsg"] = "用户不存在，请先注册"
		//l.Redirect("/login",302)
		u.TplName = "login.html"
		return
	}
	if pwd != user.Password || userName != user.Name {
		// 判断用户密码是否正确
		u.Data["errMsg"] = "用户名或密码错误！"
		//l.Redirect("/login",302)
		u.TplName = "login.html"
		return
	}
	// 4.记住用户名
	// 4.1 获取记住用户名勾选 状态
	remember := u.GetString("remember")
	// 4.2 处理复选框，不需要登录成功才存储
	if remember == "on" {
		// key value  存活时间
		u.Ctx.SetCookie("username", userName, 3600)
		u.Data["checkStatus"] = "checked"
	} else {
		// 设置存活时间为-1，不保存cookie
		u.Ctx.SetCookie("sdd", userName, -1)
		u.Data["checkStatus"] = "checked"
	}
	// 4.3 设置session，用于用户名相关操作
	u.SetSession("username", userName)

	// 5.跳转指定界面
	u.Redirect("/article/index", 302)
}

// LogOut 退出实现
func (u *UserController) LogOut() {
	// 删除session即 实现退出
	u.DelSession("username")
	u.Redirect("/login", 302)
}
