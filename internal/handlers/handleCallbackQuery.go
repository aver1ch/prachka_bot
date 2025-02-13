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

	if callbackQuery.Data == "Сушка" {
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

	if callbackQuery.Data == "Стиралка 1" || callbackQuery.Data == "Стиралка 2" || callbackQuery.Data == "Стиралка 3" {
		err := HandleLaundryButton(callbackQuery, bot, callbackQuery.Data) // нахуя тут callbackQuery.Data, если я передаю объект..
		if err != nil {
			return fmt.Errorf("%w:%w", err, errs.ErrCallbackQuery)
		}
		return nil
	}

	if callbackQuery.Data == "orderDry" {
		err := HandleOrderDryButton(callbackQuery, bot, callbackQuery.Data)
		if err != nil {
			return fmt.Errorf("%w:%w", err, errs.ErrCallbackQuery)
		}
	}

	if callbackQuery.Data == "orderLaundry+Стиралка 1" || callbackQuery.Data == "orderLaundry+Стиралка 2" || callbackQuery.Data == "orderLaundry+Стиралка 3" {
		err := HandleOrderLaundryButton(callbackQuery, bot, callbackQuery.Data)
		if err != nil {
			return fmt.Errorf("%w:%w", err, errs.ErrCallbackQuery)
		}
	}

	if callbackQuery.Data == "orderLaundry+Стиралка 1+quick" || callbackQuery.Data == "orderLaundry+Стиралка 2+quick" || callbackQuery.Data == "orderLaundry+Cтиралка 3+quick" || callbackQuery.Data == "orderLaundry+Стиралка 1+long" || callbackQuery.Data == "orderLaundry+Стиралка 2+long" || callbackQuery.Data == "orderLaundry+Стиралка 3+long" {
		err := HandleOrderLaundryConfirmButton(callbackQuery, bot, callbackQuery.Data)
		if err != nil {
			return fmt.Errorf("%w:%w", err, errs.ErrCallbackQuery)
		}
	}

	if callbackQuery.Data == "сonfirm" {
		err := HandleConfirmButton(callbackQuery, bot, callbackQuery.Data)
		if err != nil {
			return fmt.Errorf("%w:%w", err, errs.ErrCallbackQuery)
		}
	}

	if callbackQuery.Data == "myOrders" {
		err := HandleMyOrdersButton(callbackQuery, bot)
		if err != nil {
			return fmt.Errorf("%w:%w", err, errs.ErrCallbackQuery)
		}
	}

	log.Printf("Необработанный callbackQuery: %s", callbackQuery.Data)
	return nil
}
