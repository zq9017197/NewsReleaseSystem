package controllers

import (
	"github.com/astaxie/beego"
	"path"
	"time"
	"strconv"
	"math/rand"
	"github.com/astaxie/beego/orm"
	"NewsReleaseSystem/models"
)

type IndexController struct {
	beego.Controller
}

/*
ORM高级查询
项目开发中对数据库的查询，一般都是指定数据库表，用高级查询的方法进行查询。ORM支持如下几种高级查询。
Limit() 有两个参数，第一个参数是指定获取几条数据，第二个参数指定从哪里获取qs.Limit(size,start)。返回值还是qs
OrderBy() 只有一个参数，参数作用是指定按照哪个字段排序，返回值是qs。
Distinct() 没有参数，返回值是qs
Count() 没有参数，返回值是查询到的条目数和错误信息
All() 把查询到的数据全部存储到指定的容器里面，只有一个参数，指定存储查询对象的存储容器
RelatedSel() 多表查询的时候使用，指定关联的数据库表，参数长度不限，关联几个表，放几个参数
Filter() 相当于SQL语句中的where，有两个参数，第一个参数是指定查询条件，第二个参数是值
还有其他很多高级查询，具体参考：https://beego.me/docs/mvc/model/query.md页面查看
*/

//后台管理页面-显示文章列表
func (this *IndexController) ShowIndex() {
	o := orm.NewOrm()
	var articles []models.Article //存储获取的所有对象
	//查询多行数据，参数是表名，返回值是 queryseter
	qs := o.QueryTable("Article")
	//在 expr 前使用减号 - 表示 DESC 的排列
	//All 获取所有数据，相当于 select * from Article
	qs.Limit(10, 0).OrderBy("-Time").All(&articles)

	beego.Info("-------------------->",articles)
	this.Data["articles"] = articles
	this.TplName = "index.html"
}

//添加文章
func (this *IndexController) GetAdd() {
	this.TplName = "add.html"
}

/*
file,header,err := this.GetFile(key string)
作用：获取上传的静态文件。
返回值file：上传的文件，上传业务结束后需要关闭。
返回值header：上传文件的文件头。
返回值err：错误信息。

err := this.SaveToFile(fromfile, tofile string)
作用：把浏览器上传的静态文件存储。
参数fromfile：视图html中对应的name。
参数tofile：服务器端存储的文件路径。
返回值err：错误信息。
*/

//添加文章
func (this *IndexController) PostAdd() {
	title := this.GetString("articleName")
	content := this.GetString("content")
	if title == "" || content == "" {
		beego.Info("文章标题或内容不允许为空。")
		this.Data["errmsg"] = "文章标题或内容不允许为空。"
		this.TplName = "add.html"
		return
	}

	//上传图片
	file, header, err := this.GetFile("uploadname")
	defer file.Close()
	if err != nil {
		beego.Info("上传文件失败：", err)
		this.Data["errmsg"] = "上传文件失败。"
		this.TplName = "add.html"
		return
	}
	name := header.Filename //文件名
	ext := path.Ext(name)   //获取文件后缀名
	size := header.Size     //文件大小
	if ext != ".jpg" && ext != ".png" && ext != ".jpeg" {
		beego.Info("上传文件格式不正确。")
		this.Data["errmsg"] = "上传文件格式不正确。"
		this.TplName = "add.html"
		return
	} else if size > 500000 {
		beego.Info("上传文件大小超过限制。")
		this.Data["errmsg"] = "上传文件大小超过限制。"
		this.TplName = "add.html"
		return
	}
	//处理文件重名：时间 + 6位随机数 + 文件后缀名
	//"2006-01-02 15:04:05" 这个时间格式是固定值
	now := time.Now().Format("20060102150405")
	name = now + "_" + strconv.Itoa(int(rand.Int31n(1000000))) + ext
	beego.Info("长传文件名为：", name)
	path := beego.AppConfig.String("fileUpload") + name
	err = this.SaveToFile("uploadname", path)
	if err != nil {
		beego.Info("保存文件失败：", err)
		this.Data["errmsg"] = "保存文件失败。"
		this.TplName = "add.html"
		return
	}

	//保存文章数据
	o := orm.NewOrm()
	article := models.Article{}
	article.Title = title
	article.Content = content
	article.Image = path
	_, err = o.Insert(&article)
	if err != nil {
		beego.Info("保存文章失败：", err)
		this.Data["errmsg"] = "保存文章失败。"
		this.TplName = "add.html"
		return
	}
	this.Redirect("/index", 302)
}

//添加分类
func (this *IndexController) GetAddType() {
	this.TplName = "addType.html"
}

//添加分类
func (this *IndexController) PostAddType() {
	this.TplName = "addType.html"
}
