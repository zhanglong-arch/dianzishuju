package controllers

import (
	"BeegoDemo/models"
	"github.com/astaxie/beego"
	"strings"
)

type LoginController struct {
	beego.Controller
}

/**
 *直接访问login.html页面的请求
 */
func (l *LoginController) Get(){
	//设置login.html为模板文件
	//tpl: templ
	l.TplName = "login.html"
}
/**
 * 用户登录接口
 */
 func (l *LoginController) Post(){
 	var user models.User
 	err := l.ParseForm(&user)
 	if err != nil{
 		l.Ctx.WriteString("数据发生错误，请重试！")
		return
	}
	//查询数据库的用户信息
	u, err := user.QueryUser()
	if err != nil{
		l.Ctx.WriteString("用户数据查询失败，请重试！")
		return
	}
	//trim：修剪 冬青：卫矛
	//编程中：trim是将字符串中的两端的空格去掉
	name := strings.TrimSpace(u.Name)
	card := strings.TrimSpace(u.Card)
	sex := strings.TrimSpace(u.Sex)
	if name == "" || card == "" || sex == "" {
		//直接跳转到实名认证信息页面
		l.Data["Phone"] = u.Phone
		l.TplName = "user_kyc.html"
		return
	}

	//登录成功，跳转项目核心功能页面（home.html）
	l.Data["Phone"] = u.Phone
	l.TplName = "home.html"//{{.Phone}}
 }
