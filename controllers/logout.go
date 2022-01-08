package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type LogoutController struct {
	beego.Controller
}

func (c *LogoutController) Get() {
	c.DelSession("token")
	c.Ctx.Redirect(302, "/")
}
