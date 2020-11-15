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
		unix := time.Now().Format("2006_01_02 15-04-05") + ext
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
	a.TplName = "index.html"
}

// ShowContent 展示详情页面
func (a *ArticleController) ShowContent() {
	a.TplName = "content.html"
}