package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"goffredo/auth"
)

type CallbackController struct {
	beego.Controller
}

func (c *CallbackController) Get() {
	code := c.GetString("code")
	state := c.GetString("state")

	if state == "" {
		c.Ctx.Redirect(302, "/dashboard")
		return
	}

	if c.GetSession("state").(string) != state {
		c.Ctx.Abort(500, "Invalid state")
		return
	}

	tok, err := auth.UserConf.Exchange(auth.Ctx, code)
	if err != nil {
		c.Ctx.Abort(500, "Error")
		return
	}

	c.SetSession("token", tok.AccessToken)
	c.Ctx.Redirect(302, "/dashboard")
}
