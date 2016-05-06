package routers

import (
	"github.com/astaxie/beego"
	"myblog/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user/profile", &controllers.UserController{}, `get:Profile`)
	beego.Router("/api/user/profile", &controllers.UserController{}, `get:API_Profile`)
}
