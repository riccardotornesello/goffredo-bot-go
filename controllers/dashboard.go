package controllers

import (
	"fmt"
	"goffredo/auth"
	"goffredo/models"
	"path/filepath"
	"time"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/bwmarrin/discordgo"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"golang.org/x/oauth2"
)

type DashboardController struct {
	beego.Controller
}

func (c *DashboardController) Get() {
	token := c.GetSession("token")
	if token == nil {
		c.Ctx.Redirect(302, "/login")
		return
	}

	discord, err := discordgo.New("Bearer " + token.(string))
	if err != nil {
		c.Ctx.Abort(500, "Error")
		return
	}

	user, err := discord.User("@me")
	if err != nil {
		c.Ctx.Redirect(302, "/login")
		return
	}

	o := orm.NewOrm()
	var sounds []*models.Sound
	o.QueryTable(new(models.Sound)).Filter("UserId", user.ID).All(&sounds)
	soundslimit, _ := beego.AppConfig.Int("soundsLimit")

	botAddUrl := auth.BotConf.AuthCodeURL("", oauth2.SetAuthURLParam("permissions", "36700160"))

	c.Data["SoundsLimit"] = soundslimit
	c.Data["SoundsLeft"] = soundslimit - len(sounds)
	c.Data["Sounds"] = sounds
	c.Data["BotAddUrl"] = botAddUrl

	c.Layout = "layout.tpl"
	c.TplName = "dashboard.tpl"
}

func (c *DashboardController) Post() {
	token := c.GetSession("token")
	if token == nil {
		c.Ctx.Redirect(302, "/login")
		return
	}

	discord, err := discordgo.New("Bearer " + token.(string))
	if err != nil {
		c.Ctx.Abort(500, "Error")
		return
	}

	user, err := discord.User("@me")
	if err != nil {
		c.Ctx.Redirect(302, "/login")
		return
	}

	o := orm.NewOrm()
	var sounds []*models.Sound
	o.QueryTable(new(models.Sound)).Filter("UserId", user.ID).All(&sounds)
	soundslimit, _ := beego.AppConfig.Int("soundsLimit")
	if len(sounds) >= soundslimit {
		c.Ctx.Redirect(302, "/dashboard")
		return
	}

	file, _, err := c.GetFile("sound")
	if err != nil {
		c.Ctx.Abort(500, "Error")
		return
	}
	defer file.Close()
	name := c.GetString("name")

	sound := models.Sound{
		Name:   name,
		UserId: user.ID,
		Ready:  false,
	}
	id, err := o.Insert(&sound)
	if err != nil {
		c.Ctx.Abort(500, "Error")
		return
	}

	go func() {
		time.Sleep(time.Second * 10)
		soundsDirectory, _ := beego.AppConfig.String("soundsDirectory")
		soundIdStr := fmt.Sprintf("%v", id)
		outputFile := filepath.Join(soundsDirectory, soundIdStr)
		ffmpegErr := ffmpeg.Input("pipe:").
			Filter("atrim", ffmpeg.Args{}, ffmpeg.KwArgs{"start": 0, "end": 6}).
			Filter("dynaudnorm", ffmpeg.Args{}).
			Filter("volume", ffmpeg.Args{}, ffmpeg.KwArgs{"volume": "-10dB"}).
			Output(outputFile, ffmpeg.KwArgs{"format": "opus"}).
			WithInput(file).
			Run()

		if ffmpegErr != nil {
			o.Delete(&sound)
		} else {
			sound.Ready = true
			o.Update(&sound)
		}
	}()

	c.Ctx.Redirect(302, "/dashboard")
}

func (c *DashboardController) Delete() {
	token := c.GetSession("token")
	if token == nil {
		c.Ctx.Redirect(302, "/login")
		return
	}

	discord, err := discordgo.New("Bearer " + token.(string))
	if err != nil {
		c.Ctx.Abort(500, "Error")
		return
	}

	user, err := discord.User("@me")
	if err != nil {
		c.Ctx.Redirect(302, "/login")
		return
	}

	soundId := c.GetString("id")

	o := orm.NewOrm()
	o.QueryTable(new(models.Sound)).Filter("UserId", user.ID).Filter("Id", soundId).Filter("Ready", true).Delete()

	c.Data["json"] = map[string]interface{}{"status": "ok"}
	c.ServeJSON()
}
