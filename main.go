package main

import (
	"gentoo-packages-bot/handlers"
	"gentoo-packages-bot/utils"
	"log"
	"os"
	"time"

	"gopkg.in/telebot.v3"
)

func main() {
	utils.LoadEnv(".env")
	pref := telebot.Settings{
		Token: os.Getenv("TELEGRAM_BOT_TOKEN"),
		Poller: &telebot.LongPoller{
			Timeout: 10 * time.Second,
		},
	}

	b, err := telebot.NewBot(pref)

	b.Handle(telebot.OnQuery, handlers.OnQueryFunc)

	if err != nil {
		log.Fatalln(err)
	}

	b.Start()
}
