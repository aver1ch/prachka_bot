package handlers

import (
	"fmt"
	"laundryBot/internal/errs"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleCallbackQuery(callbackQuery *tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI) error {

	if callbackQuery.Data == "verify" {
		err := HandleVerifyButton(callbackQuery, bot)
		if err != nil {
			return fmt.Errorf("%w:%w", err, errs.ErrCallbackQuery)
		}
		return nil
	}

	if callbackQuery.Data == "dry" {
		err := HandleDryButton(callbackQuery, bot)
		if err != nil {
			return fmt.Errorf("%w:%w", err, errs.ErrCallbackQuery)
		}
		return nil
	}

	if callbackQuery.Data == "back" {
		err := HandleBackButton(callbackQuery, bot)
		if err != nil {
			return fmt.Errorf("%w:%w", err, errs.ErrCallbackQuery)
		}
		return nil
	}

	if callbackQuery.Data == "laundry1" || callbackQuery.Data == "laundry2" || callbackQuery.Data == "laundry3" {
		err := HandleLaundryButton(callbackQuery, bot, callbackQuery.Data)
		if err != nil {
			return fmt.Errorf("%w:%w", err, errs.ErrCallbackQuery)
		}
		return nil
	}

	log.Printf("Необработанный callbackQuery: %s", callbackQuery.Data)
	return nil
}
