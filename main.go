package main

import (
	_ "AMS/models"
	_ "AMS/routers"
	"strconv"

	"github.com/astaxie/beego"
)

func main() {
	//映射视图函数,必须放在run函数前
	_ = beego.AddFuncMap("PrePage", ShowPrePage)
	_ = beego.AddFuncMap("NextPage", ShowNextPage)
	_ = beego.AddFuncMap("autoKey", AutoKey)
	beego.Run()
}

// ShowPrePage 实现上一页
func ShowPrePage(pi int) (pre string) {
	pageIndex := pi - 1
	pre = strconv.Itoa(pageIndex)
	return
}

// ShowNextPage 实现下一页
func ShowNextPage(pi int) (next int) {
	next = pi + 1
	return
}

// AutoKey id自增
func AutoKey(key int) int {
	return key + 1
}
