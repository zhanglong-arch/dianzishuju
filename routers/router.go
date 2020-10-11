package routers

import (
	"BeegoDemo/controllers"
	"github.com/astaxie/beego"
)

/**
 * router.go文件的作用：路由功能。用于接收并分发接收到的浏览器的请求,用于匹配请求
 */
func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/user_register",&controllers.RegisterController{})
    //直接登录的页面请求接口
    beego.Router("/login.html",&controllers.LoginController{})
    beego.Router("/user_login",&controllers.LoginController{})
}
