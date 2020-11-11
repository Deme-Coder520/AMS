package main

import (
	_ "AMS/models"
	_ "AMS/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

