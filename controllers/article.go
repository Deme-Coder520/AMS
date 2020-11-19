package controllers

import (
	"AMS/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"path"
	"time"
)

type ArticleController struct {
	beego.Controller
}

// ShowIndex 展示首页
func (a *ArticleController) ShowIndex(){
	o := orm.NewOrm()
	var articles []models.Article
	_,err := o.QueryTable("Article").All(&articles)
	if err != nil {
		a.Data["code"] = "获取文章数据失败"
	}
	a.Data["articles"] = articles
	a.TplName = "index.html"
}

// ShowAdd 展示添加文章界面
func (a *ArticleController) ShowAdd() {
	a.TplName = "add.html"
}

// HandleAdd 处理添加文章业务
func (a *ArticleController) HandleAdd() {
	var filePath string
	// 1.拿到前端数据
	artName := a.GetString("artname")
	artContent := a.GetString("artcontent")

	// 2.校验数据（是否为空）
	if artName == "" || artContent == "" {
		a.Data["code"] = "Name or Content is empty"
		a.TplName = "add.html"
		return
	}
	// 3.校验文件
	file,head,err := a.GetFile("artfile")
	if file != nil {
		defer file.Close()
		if err != nil {
			a.Data["code"] = "上传图片失败"
			return
		}
		// 3.1限制文件的格式.jpg/.png/.gif
		ext := path.Ext(head.Filename)//获取文件拓展名
		if ext != ".jpg" && ext != ".png" && ext != ".gif" {
			a.Data["code"] = "文件格式不正确！"
			a.TplName = "add.html"
			return
		}
		// 3.2限制文件大小
		if head.Size > 20<<20 {
			a.Data["code"] = "文件不能大于20M"
			a.TplName = "add.html"
			return
		}
		//3.3给文件重命名
		unix := time.Now().Format("20060102_150405") + ext
		a.SaveToFile("artfile","./static/img/" + unix )//注意文件路径./开头
		filePath = "/static/img/" + unix
	}
	// 4.将数据插入数据库
	o := orm.NewOrm()
	art := models.Article{ArtName:artName,ArtContent:artContent,ArtImg:filePath}
	if _,err := o.Insert(&art);err != nil {
		beego.Info("insert date to sql fail")
		return
	}
	// 5.跳转页面（index.html）
	a.Redirect("index.html",302)
	//a.TplName = "index.html"
}

// ShowContent 展示详情页面
func (a *ArticleController) ShowContent() {
	// 1.获取文章id
	id,err  := a.GetInt("id")
	if err != nil {
		a.Data["code"] = "获取文章信息错误"
		a.Redirect("/index",302)
		return
	}
	// 2.通过id查询数据库信息
	o := orm.NewOrm()
	article := models.Article{Id:id}
	err = o.Read(&article)
	if err != nil {
		a.Data["code"] = "查询数据库信息失败"
		a.Redirect("/index",302)
		return
	}
	// 3.将数据传给视图
	a.Data["article"] = article
	// 4.跳转页面
	a.TplName = "content.html"
}

// ShowEdit 展示编辑页面
func (a *ArticleController) ShowEdit() {
	// 1.获取文章id
	id,err := a.GetInt("id")
	if err != nil {
		a.Data["code"] = "获取文章ID错误"
		a.Redirect("/index",302)
		return
	}
	// 2.根据id查询文章信息
	o := orm.NewOrm()
	article := models.Article{Id:id}
	err = o.Read(&article)
	if err != nil {
		a.Data["code"] = "查询数据信息失败"
		a.Redirect("/index",302)
		return
	}
	// 3.将文章信息传给视图
	a.Data["article"] = article
	// 4.跳转页面
	a.TplName = "edit.html"
}

// Edit 编辑文章业务处理
func (a *ArticleController) Edit() {
	// 1.获取页面数据
	id,err := a.GetInt("id")
	if err != nil {
		a.Data["code"] = "获取文章ID错误"
		a.Redirect("/index",302)
		return
	}
	name := a.GetString("artname")
	content := a.GetString("artcontent")
	// 2.校验数据
	if name == "" || content == "" {
		a.Data["code"] = "名字或内容不能为空"
		a.Redirect("/index",302)
		return
	}
	// 3.校验图片是否合法
	var filePath string
	file,head,e := a.GetFile("artfile")
	if file != nil {
		defer file.Close()
		if e != nil {
			a.Data["code"] = "上传图片失败"
			a.Redirect("/index",302)
			return
		}
		// 3.1校验图片的格式
		ext := path.Ext(head.Filename)
		if ext != ".jpg" && ext != ".png" && ext != ".gif" {
			a.Data["code"] = "图片格式不正确"
			a.Redirect("/index",302)
			return
		}
		// 3.2校验图片的大小
		if head.Size > 20<<20 {
			a.Data["code"] = "图片大小限制为20M"
			a.Redirect("/index",302)
			return
		}
		// 3.3将文件重命名
		unix := time.Now().Format("20060102_150405")+ext
		err := a.SaveToFile("artfile","./static/img/"+unix)
		if err != nil {
			beego.Info("文件保存本地失败")
			a.Redirect("/index",302)
			return
		}
		filePath = "/static/img/"+unix
	}
	// 4.更新数据库数据
	o := orm.NewOrm()
	art := models.Article{Id:id}
	err = o.Read(&art)
	if err == nil {
		if art.ArtName != name {
			art.ArtName = name
		}
		if art.ArtContent != content {
			art.ArtContent = content
		}
		if filePath != "" {
			art.ArtImg = filePath
		}
		_,err = o.Update(&art)
		if err != nil {
			a.Data["code"] = "更新数据信息失败"
			a.Redirect("/index",302)
			return
		}
		// 5.跳转页面
		a.Redirect("/index",302)
	}else{
		a.Data["code"] = "读取数据信息失败"
		a.Redirect("/index",302)
		return
	}
}

// Delete 删除业务处理
func (a *ArticleController) Delete(){
	// 1.获取文章id
	id, _ := a.GetInt("id")
	// 2.查询出对应数据并删除
	o := orm.NewOrm()
	article := models.Article{Id:id}
	err := o.Read(&article)
	if err != nil {
		a.Data["code"] = "获取文章信息失败"
		a.Redirect("/index",302)
		return
	}
	_,err = o.Delete(&article)
	if err != nil {
		a.Data["code"] = "获取文章信息失败"
	}
	// 3.跳转列表页
	a.Redirect("/index",302)
}