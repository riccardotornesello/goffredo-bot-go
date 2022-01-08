package controllers

import (
	"context"

	beego "github.com/beego/beego/v2/server/web"
	"golang.org/x/oauth2"
)

type CallbackController struct {
	beego.Controller
}

func (c *CallbackController) Get() {
	code := c.GetString("code")

	clientid, _ := beego.AppConfig.String("discord::clientid")
	clientsecret, _ := beego.AppConfig.String("discord::clientsecret")
	redirecturl, _ := beego.AppConfig.String("discord::redirecturl")

	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     clientid,
		ClientSecret: clientsecret,
		RedirectURL:  redirecturl,
		Scopes:       []string{"identify", "email", "guilds"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://discord.com/api/oauth2/authorize",
			TokenURL: "https://discord.com/api/oauth2/token",
		},
	}

	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		c.Ctx.Abort(500, "Error")
		return
	}

	c.SetSession("token", tok.AccessToken)
	c.Ctx.Redirect(302, "/dashboard")
}
