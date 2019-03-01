package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"NewsReleaseSystem/models"
)

type RegController struct {
	beego.Controller
}

//注册
func (this *RegController) GetReg() {
	this.TplName = "register.html"
}

//注册
func (this *RegController) PostReg() {
	name := this.GetString("userName")
	passwd := this.GetString("password")

	if name == "" || passwd == "" {
		beego.Info("用户名或密码不能为空！")
		this.TplName = "register.html"
	} else {
		o := orm.NewOrm()
		user := models.User{}
		user.UserName = name
		user.Passwd = passwd
		_, err := o.Insert(&user)
		if err != nil {
			beego.Error("用户【", name, "】注册失败：", err)
		}
		beego.Info("用户【", name, "】注册成功。")

		//this.Ctx.WriteString("用户【" + name + "】注册成功。")
		//渲染（转发）就是控制器把一些数据传递给视图，然后视图用这些输出组织成html界面。
		// 所以不会再给浏览器发请求，是服务器自己的行为，所以浏览器的地址栏不会改变，但是显示的页面可能会发生变化。
		//this.TplName = "login.html"
		//重定向浏览器地址栏中的 url 会发生变化。浏览中最终得到的页面是最后这个 重定向的url 请求后的页面。
		this.Redirect("/", 302) //重定向到登录页面
	}
}

/*
beego里面页面跳转的方式有两种，一种是重定向，一种是渲染。
重定向用到的方法是this.Redirect() 函数，有两个参数，第一个参数是请求路径，第二个参数是http状态码。
状态码一共分为五类：
 1xx : 服务端已经接收到了客户端请求，客户端应当继续发送请求 。常见的请求：100
 2xx :请求已成功 （已实现）常见的请求：200
 3xx :请求的资源转换路径了，请求被跳转。常见的请求：300，302
 4xx :客户端请求失败。常见的请求：404
 5xx :服务器端错误。常见的请求：500
*/
