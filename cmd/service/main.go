package main

import (
	"laundryBot/internal/handlers"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("7633580945:AAGSE71kwZY7zNLn6HpO-eNDCj_ViX1pCBk")

	if err != nil {
		log.Panic("Bot is not started: %w\n", err)
	}

	bot.Debug = true

	log.Printf("Bot started\nAutorised as %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates := bot.GetUpdatesChan(u)

	err = handlers.HandleStartButton(updates, bot)
	if err != nil {
		log.Panic(err)
	}

}
