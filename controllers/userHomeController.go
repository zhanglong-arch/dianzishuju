package controllers

import "github.com/astaxie/beego"

type HomeController struct {
	beego.Controller
}

func (l *HomeController) Get(){
	l.TplName = "home.html"
}
