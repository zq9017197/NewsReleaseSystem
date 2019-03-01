package main

import (
	_ "NewsReleaseSystem/routers"
	"github.com/astaxie/beego"
	_ "NewsReleaseSystem/models"
)

func main() {
	beego.Run()
}
