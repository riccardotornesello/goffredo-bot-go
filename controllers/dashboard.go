package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/bwmarrin/discordgo"
)

type DashboardController struct {
	beego.Controller
}

func (c *DashboardController) Get() {
	token := c.GetSession("token").(string)

	discord, err := discordgo.New("Bearer " + token)
	if err != nil {
		c.Ctx.Abort(500, "Error")
		return
	}

	user, err := discord.User("@me")
	if err != nil {
		c.Ctx.Abort(500, "Error")
		return
	}

	c.Data["Username"] = user.Username
	c.TplName = "dashboard.tpl"
}
