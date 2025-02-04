package handlers

import (
	"laundryBot/internal/send"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var UserState = make(map[int64]string)

func HandleStartButton(update tgbotapi.Update, bot *tgbotapi.BotAPI) error {
	if update.Message != nil && update.Message.IsCommand() && update.Message.Command() == "start" {
		chatID := update.Message.Chat.ID
		userName := update.Message.From.UserName

		err := send.SendStartMessage(chatID, userName, bot)
		if err != nil {
			return err
		}
	}
	return nil
}
