package routers

import (
	"goffredo/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/callback", &controllers.CallbackController{})
	beego.Router("/dashboard", &controllers.DashboardController{})
}
