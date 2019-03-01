package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"NewsReleaseSystem/models"
)

type MainController struct {
	beego.Controller
}

//登录
func (c *MainController) Get() {
	c.TplName = "login.html"
}

/*
this.GetString()：获取字符串类型值
this.GetInt()：获取整型值
this.GetFloat：获取浮点型值
...
this.GetFile()：获取上传的文件
作用：接收前端传递过来的数据，不管是get请求还是post请求，都能接收。
参数: 是传递数据的key值，一般情况下是form表单中标签的name属性值
返回值：根据返回类型不同，返回值也不一样，最常用的GetString()只有一个返回值，如果没有取到值就返回空字符串，其他几个函数会返回一个错误类型。
获取的值一般是标签里面的value属性值。
*/

//登录
func (c *MainController) Post() {
	name := c.GetString("userName")
	passwd := c.GetString("password")
	if name == "" || passwd == "" {
		beego.Info("用户名或密码不能为空！")
		c.TplName = "register.html"
		return
	}

	o := orm.NewOrm()
	user := models.User{}
	user.UserName = name
	err := o.Read(&user, "UserName")
	if err != nil {
		beego.Info("登录失败：", err)
		c.TplName = "login.html"
		return
	}

	if user.Passwd != passwd {
		beego.Info("登录失败，密码错误。")
		c.TplName = "login.html"
		return
	}

	beego.Info("登录成功。")
	c.Data["name"] = user.UserName
	c.Redirect("/index", 302) //重定向到后台管理页面
}
