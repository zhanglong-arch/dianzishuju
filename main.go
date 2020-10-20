package main

import (
	"BeegoDemo/blockchain"
	"BeegoDemo/db_mysql"
	_ "BeegoDemo/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {

	//1、生成第一个区块
	block := blockchain.NewBlock(0, []byte{}, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})

	fmt.Println(block)
	fmt.Printf("区块的Hash值：%x",block.Hash)
	return

	db_mysql.ConnerDB()
	//静态资源路径设置
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	beego.Run()
}

