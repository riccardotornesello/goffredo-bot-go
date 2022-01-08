package main

import (
	_ "goffredo/db"
	_ "goffredo/routers"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"goffredo/bot"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	// Run bot
	var wg sync.WaitGroup
	wg.Add(1)
	go func() { bot.Run(&wg) }()

	// Run web server
	go func() { beego.Run() }()

	// Wait for interrupt
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	wg.Done()

	time.Sleep(1 * time.Second)
}
