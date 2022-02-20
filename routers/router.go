package routers

import (
	"goffredo/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.SetStaticPath("/static","static")
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LogoutController{})
	beego.Router("/callback", &controllers.CallbackController{})
	beego.Router("/dashboard", &controllers.DashboardController{})
}
