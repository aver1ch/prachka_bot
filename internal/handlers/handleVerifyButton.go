package handlers

import (
	"fmt"
	"laundryBot/internal/db"
	"laundryBot/internal/errs"
	"laundryBot/internal/processing"
	"laundryBot/internal/send"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleVerifyButton(callbackQuery *tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI) error {
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
		err := send.SendAlreadyVerificatedMessage(chatID, userName, bot)
		if err != nil {
			return fmt.Errorf("%w:%w", err, errs.ErrSendMessage)
		}

		err = send.SendChooseMenu(chatID, userName, bot)
		if err != nil {
			return fmt.Errorf("%w:%w", err, errs.ErrSendMessage)
		}

		return nil
	}

	UserState[chatID] = "waiting_for_room_number"

	err = send.SendNumberOfRoomRequestMessage(chatID, userName, bot)
	if err != nil {
		return err
	}

	log.Printf("Пользователь %s (chatID: %d) нажал на кнопку 'Верифицироваться'", userName, chatID)
	return nil
}

func HandleRoomNumberMessage(update tgbotapi.Update, bot *tgbotapi.BotAPI) error {

	chatID := update.Message.Chat.ID

	if UserState[chatID] != "waiting_for_room_number" {
		return nil
	}

	roomNumber := update.Message.Text
	username := update.Message.From.UserName

	err := processing.ProcessRoomNumber(roomNumber)
	if err != nil {
		sendErr := send.SendErrorNumberOfRoomRequestMessage(chatID, username, bot)
		if sendErr != nil {
			return sendErr
		}
		return err
	}

	dbConn, err := db.ConnectToDB()
	if err != nil {
		return err
	}
	defer dbConn.Close()

	err = db.InsertUserToDB(dbConn, username, roomNumber, true)
	if err != nil {
		return err
	}

	err = send.SendSuccessVerificationMessage(chatID, username, bot)
	if err != nil {
		return err
	}

	UserState[chatID] = ""

	log.Printf("Пользователь %s (chatID: %d) успешно верифицирован", username, chatID)

	send.SendChooseMenu(chatID, username, bot)

	return nil
}
