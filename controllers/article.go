package controllers

import (
	"AMS/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"math"
	"path"
	"strconv"
	"time"
)

type ArticleController struct {
	beego.Controller
}

// ShowIndex 展示首页并实现分页功能
func (a *ArticleController) ShowIndex(){
	// 1.查询所有文章数据
	o := orm.NewOrm()
	var articles []models.Article
	querySeter := o.QueryTable("Article")
	/*_,err := querySeter.All(&articles)*/
	count,_ := querySeter.Count()
	// 2.设置每一页显示的数量，从而得到总的页数
	var pageSize =2
	pageCount := math.Ceil(float64(count)/float64(pageSize))// 向上取整，显示的页面不会出现小数
	// 3.首页和末页
	pi := a.GetString("pi")
	pageIndex,err := strconv.Atoi(pi)
	if err != nil {
		pageIndex = 1// 首页没有传pageIndex的值，防止默认pageIndex为0
	}
	// 3.1每一页显示的个数
	stat := pageSize*(pageIndex-1)
	_,err = querySeter.Limit(pageSize,stat).All(&articles)
	if err != nil {
		beego.Info("获取文章数据失败")
		a.Redirect("/index",302)
		return
	}
	// 4.上一页和下一页限制(视图函数)
	var isFirstPage = false
	var isLastPage = false
	if pageIndex == 1 {
		isFirstPage = true
	}
	if pageIndex == int(pageCount) {
		isLastPage = true
	}

	a.Data["count"] = count
	a.Data["pageCount"] = pageCount
	a.Data["pageIndex"] = pageIndex
	a.Data["isFirstPage"] = isFirstPage
	a.Data["isLastPage"] = isLastPage
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
		beego.Info("文章名或内容不能为空")
		a.Redirect("/addArticle",302)
		return
	}
	// 3.校验文件
	file,head,err := a.GetFile("artfile")
	if file != nil {
		defer file.Close()
		if err != nil {
			beego.Info("上传文件失败")
			a.Redirect("/addArticle",302)
			return
		}
		// 3.1限制文件的格式.jpg/.png/.gif
		ext := path.Ext(head.Filename)//获取文件拓展名
		if ext != ".jpg" && ext != ".png" && ext != ".gif" {
			beego.Info("文件格式不正确！")
			a.Redirect("/addArticle",302)
			return
		}
		// 3.2限制文件大小
		if head.Size > 20<<20 {
			beego.Info("文件不能大于20M")
			a.Redirect("/addArticle",302)
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
	_,err = o.Insert(&art)
	if err != nil {
		beego.Info("添加图片至数据库失败")
		a.Redirect("/addArticle",302)
		return
	}
	// 5.跳转页面（index.html）
	a.Redirect("/index",302)
	//a.TplName = "index.html"
}

// ShowContent 展示详情页面
func (a *ArticleController) ShowContent() {
	// 1.获取文章id
	id,err  := a.GetInt("id")
	if err != nil {
		beego.Info("获取文章信息错误")
		a.Redirect("/index",302)
		return
	}
	// 2.通过id查询数据库信息
	o := orm.NewOrm()
	article := models.Article{Id:id}
	err = o.Read(&article)
	if err != nil {
		beego.Info("查询数据库信息失败")
		a.Redirect("/index",302)
		return
	}
	// 3.将数据传给视图
	a.Data["article"] = article
	// 4.跳转页面
	a.Redirect("/content",302)
}

// ShowEdit 展示编辑页面
func (a *ArticleController) ShowEdit() {
	// 1.获取文章id
	id,err := a.GetInt("id")
	if err != nil {
		beego.Info("获取文章ID错误")
		a.Redirect("/index",302)
		return
	}
	// 2.根据id查询文章信息
	o := orm.NewOrm()
	article := models.Article{Id:id}
	err = o.Read(&article)
	if err != nil {
		beego.Info("查询数据信息失败")
		a.Redirect("/index",302)
		return
	}
	// 3.将文章信息传给视图
	a.Data["article"] = article
	// 4.跳转页面
	a.Redirect("/edit",302)
}

// Edit 编辑文章业务处理
func (a *ArticleController) Edit() {
	// 1.获取页面数据
	id,err := a.GetInt("id")
	if err != nil {
		beego.Info("获取文章ID错误")
		a.Redirect("/index",302)
		return
	}
	name := a.GetString("artname")
	content := a.GetString("artcontent")
	// 2.校验数据
	if name == "" || content == "" {
		beego.Info("名字或内容不能为空")
		a.Redirect("/index",302)
		return
	}
	// 3.校验图片是否合法
	var filePath string
	file,head,e := a.GetFile("artfile")
	if file != nil {
		defer file.Close()
		if e != nil {
			beego.Info("上传图片失败")
			a.Redirect("/index",302)
			return
		}
		// 3.1校验图片的格式
		ext := path.Ext(head.Filename)
		if ext != ".jpg" && ext != ".png" && ext != ".gif" {
			beego.Info("图片格式不正确")
			a.Redirect("/index",302)
			return
		}
		// 3.2校验图片的大小
		if head.Size > 20<<20 {
			beego.Info("图片大小限制为20M")
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
			beego.Info("更新数据信息失败")
			a.Redirect("/index",302)
			return
		}
		// 5.跳转页面
		a.Redirect("/index",302)
	}else{
		beego.Info("读取数据信息失败")
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
		beego.Info("获取文章信息失败")
		a.Redirect("/index",302)
		return
	}
	_,err = o.Delete(&article)
	if err != nil {
		beego.Info("获取文章信息失败")
		a.Redirect("/index",302)
		return
	}
	// 3.跳转列表页
	a.Redirect("/index",302)
}