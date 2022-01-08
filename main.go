package main

import (
	_ "goffredo/db"
	_ "goffredo/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}
