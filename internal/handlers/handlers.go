package handlers

import (
	"laundryBot/internal/send"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleStartButton(updates tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI) error {
	for update := range updates {
		if update.Message != nil && update.Message.IsCommand() && update.Message.Command() == "start" {
			chatID := update.Message.Chat.ID
			userName := update.Message.From.FirstName
			err := send.SendStartMessage(chatID, userName, bot)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
