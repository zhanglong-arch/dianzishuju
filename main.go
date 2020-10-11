package main

import (
	"BeegoDemo/db_mysql"
	_ "BeegoDemo/routers"
	"github.com/astaxie/beego"
)

func main() {
	db_mysql.ConnerDB()
	//静态资源路径设置
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	beego.Run()
}

