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
	beego.Router("/index", &controllers.IndexController{}, "get:ShowIndex")
	//添加文章
	beego.Router("/add", &controllers.IndexController{}, "get:GetAdd;post:PostAdd")
	//添加分类
	beego.Router("/addType", &controllers.IndexController{}, "get:GetAddType;post:PostAddType")
	//后台管理页面-文章列表-详情
	beego.Router("/articleDetail", &controllers.IndexController{}, "get:ArticleDetail")
	//后台管理页面-文章列表-删除
	beego.Router("/articleDelete", &controllers.IndexController{}, "get:ArticleDelete")
	//后台管理页面-文章列表-编辑
	beego.Router("/articleModify", &controllers.IndexController{}, "get:GetArticleModify;post:PostArticleModify")
}
