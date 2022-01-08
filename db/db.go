package db

import (
	"goffredo/models"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	orm.RegisterModel(new(models.Sound))
	orm.RegisterDataBase("default", "sqlite3", "dev.db")
	orm.RunSyncdb("default", false, true)
}
