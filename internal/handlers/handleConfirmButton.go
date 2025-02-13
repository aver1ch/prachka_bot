package handlers

import (
	"fmt"
	"laundryBot/internal/db"
	"laundryBot/internal/errs"
	"laundryBot/internal/send"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleConfirmButton(callbackQuery *tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI, service string) error {
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
		err := send.SendConfirmMessage(chatID, userName, bot)
		log.Printf("Пользователь %s (chatID: %d) нажал на кнопку 'Подтвердить бронь'", userName, chatID)
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
