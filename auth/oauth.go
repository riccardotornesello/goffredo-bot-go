package auth

import (
	"context"

	beego "github.com/beego/beego/v2/server/web"
	"golang.org/x/oauth2"
)

var Ctx context.Context
var UserConf *oauth2.Config
var BotConf *oauth2.Config

func init() {
	clientid, _ := beego.AppConfig.String("discord::clientid")
	clientsecret, _ := beego.AppConfig.String("discord::clientsecret")
	redirecturl, _ := beego.AppConfig.String("discord::redirecturl")

	Ctx = context.Background()
	UserConf = &oauth2.Config{
		ClientID:     clientid,
		ClientSecret: clientsecret,
		RedirectURL:  redirecturl,
		Scopes:       []string{"identify", "email", "guilds"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://discord.com/api/oauth2/authorize",
			TokenURL: "https://discord.com/api/oauth2/token",
		},
	}
	BotConf = &oauth2.Config{
		ClientID:     clientid,
		ClientSecret: clientsecret,
		RedirectURL:  redirecturl,
		Scopes:       []string{"bot"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://discord.com/api/oauth2/authorize",
			TokenURL: "https://discord.com/api/oauth2/token",
		},
	}
}
