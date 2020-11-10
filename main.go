package main

import (
	"BeegoDemo/blockchain"
	"BeegoDemo/db_mysql"
	_ "BeegoDemo/routers"
	"github.com/astaxie/beego"
)

func main() {

	//先准备一条区块链
	blockchain.NewBlockChain()

	//1、链接数据库
	db_mysql.ConnerDB()
	//静态资源路径设置
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")

	//3、允许
	beego.Run() //启动端口监听：阻塞
}

