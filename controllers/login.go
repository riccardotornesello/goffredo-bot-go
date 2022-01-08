package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	clientid, _ := beego.AppConfig.String("discord::clientid")
	clientsecret, _ := beego.AppConfig.String("discord::clientsecret")
	redirecturl, _ := beego.AppConfig.String("discord::redirecturl")

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

	state := uuid.New().String()
	url := conf.AuthCodeURL(state, oauth2.AccessTypeOffline)
	c.SetSession("state", state)

	c.Ctx.Redirect(302, url)
}
