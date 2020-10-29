package main

import (
	"BeegoDemo/blockchain"
	"BeegoDemo/db_mysql"
	_ "BeegoDemo/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	//1、实例化一个区块链实例
	bc := blockchain.NewBlockChain()
	fmt.Println(bc)
	return
	//fmt.Printf("最新区块的Hash值：%x\n",bc.LastHash)
	////数据被保存早height为1的区块当中
	//fmt.Println(bc)
	//return
	//block, err := bc.SaveData([]byte("这里存储上链的数据信息"))
	//fmt.Println(block)
	////return
	//if err != nil{
	//	fmt.Println(err.Error())
	//	return
	//}

	//block1 := bc.QueryBlockByHeight(1)
	//if block1 == nil{
	//	fmt.Println("抱歉，输入有误")
	//	return
	//}
	//fmt.Println("区块的高度是：",block1.Height)
	//fmt.Println("区块存的信息是：",string(block1.Data))
	//blocks := bc.QueryAllBlocks()
	//if len(blocks) == 0 {
	//	fmt.Println("暂未查询到区块数据")
	//	return
	//}
	//for _, block := range blocks{
	//	fmt.Println("高度：%d,哈希：%x,Prev哈希：%x\n",block.Height,block.Hash,block.PrevHash)
	//}
	//return


	db_mysql.ConnerDB()
	//静态资源路径设置
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	beego.Run()
}

