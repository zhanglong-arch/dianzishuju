package controllers

import (
	"BeegoDemo/models"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

/**
 * 该方法用于处理在浏览器直接请求用户注册页面
 */
func (r *RegisterController) Get(){
	r.TplName = "register.html"
}

/**
 * 该方法用于处理用户注册的表单提交请求
 */
func (r *RegisterController) Post(){
	//1、解析请求数据
	var user models.User
	err := r.ParseForm(&user)
	if err != nil {
		//返回错误信息给浏览器，提示用户
		r.TplName="error.html"
		return
	}
	//2、保存用户信息到数据库
	_, err = user.SaveUser()
	//3、返回前端结果（成功跳登录页面，失败弹出错误信息）
	if err != nil {
		r.TplName="error.html"
		return
	}
	//用户注册成功
	r.TplName = "login.html"
}
