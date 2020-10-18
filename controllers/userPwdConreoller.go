package controllers

import "github.com/astaxie/beego"

type PwdController struct {
	beego.Controller
}
func (l *PwdController) Get(){
	l.TplName="password.html"
}