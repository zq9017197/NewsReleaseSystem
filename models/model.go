package models

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
	"github.com/astaxie/beego/orm"
)

//用户表
type User struct {
	Id       int
	UserName string `orm:"size(50)"` //用户名
	Passwd   string `orm:"size(50)"` //密码
}

/*
我们以前在数据库中创建表的时候，会给字段加很多限制属性，比如非空，长度，默认值等等，在ORM中，创建表时也可以给各个字段添加相应的限制。
要给哪个字段添加属性，需要在这个字段后面添加 `` 括起来的内容，格式为orm:"限制条件" 。
 pk 设置该字段为主键
 auto 这只该字段自增，但是要求该字段必须为整型
 default(0) 设置该字段的默认值，需要注意字段类型和默认值类型一致
 size(100) 设置该字段长度为100个字节，一般用来设置字符串类型
 null 设置该字段允许为空，默认不允许为空
 unique 设置该字段全局唯一
 digits(12);decimals(4) 设置浮点数位数和精度。比如这个是说，浮点数总共12位，小数位为四位。
 auto_now 针对时间类型字段，作用是保存数据的更新时间
 auto_now_add 针对时间类型字段,作用是保存数据的添加时间
Mysql中时间类型有date和datetime两种类型，但是我们go里面只有time.time一种类型，如果项目里面要求精确的话，就需要指定类型，指定类型用的是type(date)或者type(datetime)
注意：当模型定义里没有主键时，符合int, int32, int64, uint, uint32, uint64 类型且名称为 Id 的 Field 将被视为主键，能够自增.
*/

//文章表
type Article struct {
	Id      int
	Title   string    `orm:"size(20)"`                    //标题
	Content string    `orm:"size(500)"`                   //内容
	Image   string    `orm:"size(200);null"`               //图片
	Time    time.Time `orm:"type(datetime);auto_now_add"` //发布时间
	Count   int       `orm:"default(0)"`                  //阅读量
}

func init() {
	//create database newsweb charset=utf8;
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(180.76.248.109:3306)/newsweb?charset=utf8")
	orm.RegisterModel(new(User), new(Article))
	orm.RunSyncdb("default", false, true)
}
