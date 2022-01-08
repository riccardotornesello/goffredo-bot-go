package bot

import (
	"goffredo/models"
	"math/rand"

	"github.com/beego/beego/v2/client/orm"
)

func GetSound(userId string) *models.Sound {
	o := orm.NewOrm()
	var sounds []*models.Sound
	num, err := o.QueryTable(new(models.Sound)).Filter("UserId", userId).Filter("Ready", true).All(&sounds)

	if err != nil {
		return nil
	}

	if num == 0 {
		return nil
	}

	return sounds[rand.Intn(len(sounds))]
}
