package main

import (
	"github.com/astaxie/beego"
	_ "myblog/models"
	"myblog/models/class"
	_ "myblog/routers"
)

func main() {
	class.TestORM()
	beego.Run()
}
