package handlers

import (
	"context"
	"fmt"
	"laundryBot/internal/db"
	"laundryBot/internal/errs"
	"laundryBot/internal/send"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleDryButton(callbackQuery *tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI) error {
	chatID := callbackQuery.Message.Chat.ID
	userName := callbackQuery.From.UserName

	dbConn, err := db.ConnectToDB()
	if err != nil {
		return err
	}
	defer dbConn.Disconnect(context.Background())

	usersCollection := dbConn.Database("dorm1").Collection("users")

	isAuthorized, err := db.GetIsAuthorisedFromDB(usersCollection, userName)
	if err != nil {
		return fmt.Errorf("%w:%w (%v, %v)", err, errs.ErrPullingDataFromDB, chatID, userName)
	}

	if isAuthorized {
		err := send.SendInfoByService(chatID, userName, bot, "dry")
		log.Printf("Пользователь %s (chatID: %d) нажал на кнопку 'Cушилка'", userName, chatID)
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
