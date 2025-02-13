package handlers

import (
	"fmt"
	"laundryBot/internal/db"
	"laundryBot/internal/errs"
	"laundryBot/internal/send"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleOrderLaundryButton(callbackQuery *tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI, mode string) error {
	chatID := callbackQuery.Message.Chat.ID
	userName := callbackQuery.From.UserName

	dbConn, err := db.ConnectToDB()
	if err != nil {
		return err
	}
	defer dbConn.Close()

	isAuthorized, err := db.GetIsAuthorisedFromDB(dbConn, userName)
	if err != nil {
		return fmt.Errorf("%w: %w", err, errs.ErrAuthorizationError)
	}

	if isAuthorized {
		err := send.SendRequestOfLaundryMode(chatID, userName, bot, mode)
		log.Printf("Пользователь %s (chatID: %d) нажал на кнопку 'Cтиралка'", userName, chatID)
		if err != nil {
			return fmt.Errorf("%w:%w", err, errs.ErrAlreadyAutorized)
		}
	} else {
		err = send.SendVerificationError(chatID, userName, bot)
		if err != nil {
			return err
		}
	}
	return nil
}

func HandleOrderLaundryConfirmButton(callbackQuery *tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI, mode string) error {
	chatID := callbackQuery.Message.Chat.ID
	userName := callbackQuery.From.UserName

	dbConn, err := db.ConnectToDB()
	if err != nil {
		return err
	}
	defer dbConn.Close()

	isAuthorized, err := db.GetIsAuthorisedFromDB(dbConn, userName)
	if err != nil {
		return fmt.Errorf("%w: %w", err, errs.ErrAuthorizationError)
	}

	if isAuthorized {
		err := send.SendRequestOfOrderConfirmation(chatID, userName, bot, mode)
		log.Printf("Пользователь %s (chatID: %d) нажал на кнопку 'Подтвердить бронь' по %s", userName, chatID, mode)
		if err != nil {
			return fmt.Errorf("%w:%w", err, errs.ErrAlreadyAutorized)
		}
	} else {
		err = send.SendVerificationError(chatID, userName, bot)
		if err != nil {
			return err
		}
	}
	return nil
}
