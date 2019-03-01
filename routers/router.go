package routers

import (
	"NewsReleaseSystem/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//登录页面
	beego.Router("/", &controllers.MainController{})
	//注册页面
	beego.Router("/register", &controllers.RegController{}, "get:GetReg;post:PostReg")
	//后台管理页面-文章列表
	beego.Router("/index", &controllers.IndexController{}, "get,post:ShowIndex")
	//添加文章
	beego.Router("/add", &controllers.IndexController{}, "get:GetAdd;post:PostAdd")
	//添加分类
	beego.Router("/addType", &controllers.IndexController{}, "get:GetAddType;post:PostAddType")
}
