package routers

import (
	"BeegoDemo/controllers"
	"github.com/astaxie/beego"
)

/**
 * router.go文件的作用：路由功能。用于接收并分发接收到的浏览器的请求,用于匹配请求
 */
func init() {
	//注册页面
    beego.Router("/", &controllers.MainController{})
    //用户注册的接口请求
    beego.Router("/user_register",&controllers.RegisterController{})
    //直接登录的页面请求接口
    beego.Router("/login.html",&controllers.LoginController{})
    //用户登录请求接口
    beego.Router("/user_login",&controllers.LoginController{})
    //请求用户注册的页面
    beego.Router("/zhuce_register",&controllers.RegisterController{})
    //存证确权跳转页面
    beego.Router("/list_record.html",&controllers.CunZhengQueQuan{})
    //忘记密码页面接口请求
    beego.Router("/pwd_password",&controllers.PwdController{})
    //文件上传接口
    beego.Router("/upload",&controllers.HomeController{})
    //在认证数据列表页面，点击新增认证按钮，跳转“新增页面”
    beego.Router("/upload_file.html",&controllers.HomeController{})
    //查看认证数据的证书(cert_detail.html)
    beego.Router("/cert_detail.html",&controllers.CerDetailController{})
    //浏览器中发起的连接跳转：用户实名认证
    beego.Router("/user_kyc.html",&controllers.UserKycController{})
    //用户实名认证功能接口
    beego.Router("/user_kyc",&controllers.UserKycController{})
    //短信验证登录页面
    beego.Router("/login_sms.html",&controllers.SmsLoginController{})
    //发送验证码短信息
    beego.Router("/send_sms",&controllers.SendSmsController{})
    //调用登录接口，执行手机号和验证码登录
    beego.Router("/login_sms",&controllers.SmsLoginController{})
}
