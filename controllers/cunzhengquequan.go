package controllers

import "github.com/astaxie/beego"

type CunZhengQueQuan struct {
	beego.Controller
}

func (c *CunZhengQueQuan) Get(){
	c.TplName = "list_record.html"
}
