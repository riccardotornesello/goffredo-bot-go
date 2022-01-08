package controllers

import (
	"goffredo/models"

	"github.com/beego/beego/v2/client/orm"
)

type BotController struct{}

func (c *BotController) GetSound(userId string) *models.Sound {
	o := orm.NewOrm()
	var sound *models.Sound
	err := o.QueryTable(new(models.Sound)).Filter("UserId", userId).Filter("Ready", true).OrderBy("RANDOM()").Limit(1).One(&sound)

	if err == orm.ErrNoRows {
		return nil
	}

	return sound
}
