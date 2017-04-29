package main

import (
	_ "static-server/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

