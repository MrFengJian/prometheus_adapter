package main

import (
	_ "prometheus_adapter/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.AppPath="./conf/app.conf"
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
