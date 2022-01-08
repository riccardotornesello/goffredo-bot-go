package controllers

import (
	"fmt"
	"goffredo/models"
	"path/filepath"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/bwmarrin/discordgo"
	ffmpeg "github.com/u2takey/ffmpeg-go"
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

	o := orm.NewOrm()
	var sounds []*models.Sound
	o.QueryTable(new(models.Sound)).Filter("UserId", user.ID).All(&sounds)

	soundslimit, _ := beego.AppConfig.Int("soundsLimit")
	c.Data["SoundsLimit"] = soundslimit
	c.Data["SoundsLeft"] = soundslimit - len(sounds)
	c.Data["Sounds"] = sounds

	c.Layout = "layout.tpl"
	c.TplName = "dashboard.tpl"
}

func (c *DashboardController) Post() {
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
	o := orm.NewOrm()
	id, err := o.Insert(&sound)
	if err != nil {
		c.Ctx.Abort(500, "Error")
		return
	}

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
		c.Ctx.Abort(500, ffmpegErr.Error())
		return
	}

	sound.Ready = true
	o.Update(&sound)

	c.Ctx.Redirect(302, "/dashboard")
}
