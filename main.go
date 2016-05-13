package main

import (
	"github.com/crimesbot/bot/settings"
	"github.com/crimesbot/bot/web"
)

func main() {
	// Start this webserver just to never puts this instance idle
	go web.StartServer(settings.IP, settings.Port)
	// Creating and starting a new bot
	NewBot(settings.Token).Run()
}
