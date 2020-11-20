package main

import (
	_ "AMS/models"
	_ "AMS/routers"
	"github.com/astaxie/beego"
	"strconv"
)

func main() {
	//映射视图函数,必须放在run函数前
	_=beego.AddFuncMap("PrePage",ShowPrePage)
	_=beego.AddFuncMap("NextPage",ShowNextPage)
	beego.Run()
}

// ShowPrePage 实现上一页
func ShowPrePage(pi int)(pre string){
	pageIndex := pi -1
	pre = strconv.Itoa(pageIndex)
	return
}

// ShowPrePage 实现下一页
func ShowNextPage(pi int)(next string){
	pageIndex := pi + 1
	next = strconv.Itoa(pageIndex)
	return
}

