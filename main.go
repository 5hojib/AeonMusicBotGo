package main

import (
	"log"

	"github.com/amarnathcjd/gogram/telegram"
	"AeonMusisBotGo/config"
	"AeonMusisBotGo/modules"
)

var (
	user *telegram.Client
	bot  *telegram.Client
)

func main() {
	config.LoadConfig()

	var err error

	bot, err = telegram.NewClient(
		telegram.ClientConfig{
			AppID:    config.APP_ID,
			AppHash:  config.API_HASH,
			LogLevel: telegram.LogError,
			Session:  "bot.dat",
			Cache:    telegram.NewCache("bot_cache"),
		},
	)
    if err != nil {
        log.Fatal("Error creating bot client:", err)
    }

    if err := bot.LoginBot(config.BOT_TOKEN); err != nil {
        log.Fatal("Error logging in bot:", err)
    }

    user, err = telegram.NewClient(
        telegram.ClientConfig{
            AppID:         config.APP_ID,
            AppHash:       config.API_HASH,
            StringSession: config.USER_SESSION,
            LogLevel:      telegram.LogError,
            Session:       "user.dat",
            Cache:         telegram.NewCache("user_cache"),
        },
    )
    if err != nil {
        log.Fatal("Error creating user client:", err)
    }

    if err := user.Start(); err != nil {
        log.Fatal("Error starting user client:", err)
    }

    bot.On(
        "message:/start",
        func(message *telegram.NewMessage) error {
            _, err := message.Reply("Hello, I am a bot!")
            return err
        },
    )

    bot.On(
        "message:/check",
        func(message *telegram.NewMessage) error {
            _, err := user.SendMessage("me", "done")
            return err
        },
    )

    bot.AddInlineHandler(telegram.OnInlineQuery, modules.InlineMusicSearch)
    bot.On(telegram.OnChoosenInline, modules.ChosenInlineHandler)
    bot.Idle()
}