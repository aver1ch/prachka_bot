package handlers

import (
	"fmt"
	"laundryBot/internal/errs"
	"laundryBot/internal/send"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleBackButton(callbackQuery *tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI) error {
	chatID := callbackQuery.Message.Chat.ID
	userName := callbackQuery.From.UserName

	err := send.SendChooseMenu(chatID, userName, bot)
	if err != nil {
		return fmt.Errorf("%w:%w", err, errs.ErrSendMessage)
	}
	return nil
}
