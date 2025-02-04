package send

import (
	"fmt"
	"laundryBot/internal/errs"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendStartMessage(chatID int64, userName string, bot *tgbotapi.BotAPI) error {
	msg := tgbotapi.NewMessage(chatID, "Привет! Этот бот поможет решить проблему со стиралками, теперь тебе не нужно постоянно переписываться в чатике по поводу стиралок.\nДостаточно нажать пару кнопок в этом боте, чтобы занять стиралку или сушилку :)\n\nДля начала давай верифицируемся\n\nВерификация возможна ТОЛЬКО ОДИН РАЗ. Если вы ошиблись, то напишите админу: @averichie")
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		[]tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardButtonData("Верифицироваться", "verify"),
		},
	)

	msg.ReplyMarkup = inlineKeyboard

	_, err := bot.Send(msg)
	if err != nil {
		return fmt.Errorf("%w:%w\n", err, errs.ErrSendMessage)
	}

	log.Printf("Стартовое сообщение отправлено %s (chatID: %d)", userName, chatID)
	return nil
}

func SendNumberOfRoomRequestMessage(chatID int64, userName string, bot *tgbotapi.BotAPI) error {
	msg := tgbotapi.NewMessage(chatID, "Введи номер своей комнаты")

	_, err := bot.Send(msg)
	if err != nil {
		return fmt.Errorf("%w:%w\n", err, errs.ErrSendMessage)
	}

	log.Printf("Запрос о вводе номера комнаты отправлен %s (chatID: %d)", userName, chatID)
	return nil
}

func SendErrorNumberOfRoomRequestMessage(chatID int64, userName string, bot *tgbotapi.BotAPI) error {
	msg := tgbotapi.NewMessage(chatID, "Некорректный ввод. Попробуй ещё раз :)")

	_, err := bot.Send(msg)
	if err != nil {
		return fmt.Errorf("%w:%w\n", err, errs.ErrSendMessage)
	}

	return nil
}

func SendSuccessVerificationMessage(chatID int64, userName string, bot *tgbotapi.BotAPI) error {
	msg := tgbotapi.NewMessage(chatID, "Верификация прошла успешно")

	_, err := bot.Send(msg)
	if err != nil {
		return fmt.Errorf("%w:%w\n", err, errs.ErrSendMessage)
	}

	return nil
}

func SendAlreadyVerificatedMessage(chatID int64, userName string, bot *tgbotapi.BotAPI) error {
	msg := tgbotapi.NewMessage(chatID, "Вы уже верифицированы, выберите, что вы хотите занять")

	_, err := bot.Send(msg)
	if err != nil {
		return fmt.Errorf("%w:%w\n", err, errs.ErrSendMessage)
	}

	return nil
}

func SendVerificationError(chatID int64, userName string, bot *tgbotapi.BotAPI) error {
	msg := tgbotapi.NewMessage(chatID, "Тебе нужно верифицироваться, прежде чем пользоваться этим ботом\n\nВерификация возможна ТОЛЬКО ОДИН РАЗ. Если вы ошиблись, то напишите админу: @averichie")
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		[]tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardButtonData("Верифицироваться", "verify"),
		},
	)

	msg.ReplyMarkup = inlineKeyboard

	_, err := bot.Send(msg)
	if err != nil {
		return fmt.Errorf("%w:%w\n", err, errs.ErrSendMessage)
	}

	log.Printf("Сообщение о верификации отправлено %s (chatID: %d)", userName, chatID)
	return nil
}

func SendChooseMenu(chatID int64, userName string, bot *tgbotapi.BotAPI) error {
	msg := tgbotapi.NewMessage(chatID, "Ты будешь стираться или сушиться?")
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Сушка", "dry"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Стиралка 1", "laundry1"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Стиралка 2", "laundry2"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Стиралка 3", "laundry3"),
		),
	)
	msg.ReplyMarkup = inlineKeyboard

	_, err := bot.Send(msg)
	if err != nil {
		return fmt.Errorf("%w:%w\n", err, errs.ErrSendMessage)
	}

	log.Printf("Сообщение о выборе услуги отправлено %s (chatID: %d)", userName, chatID)
	return nil
}

// забронировать на быструю стирку или другой режим
func SendInfoByService(chatID int64, userName string, bot *tgbotapi.BotAPI, service string) error {
	var msg tgbotapi.MessageConfig
	if service == "dry" {
		info := fmt.Sprintf("информация по сушке %s", service)
		msg = tgbotapi.NewMessage(chatID, info)
		log.Printf("Сообщение о сушке отправлено %s (chatID: %d)", userName, chatID)
	} else {
		info := fmt.Sprintf("информация по стиралке %s", service)
		msg = tgbotapi.NewMessage(chatID, info)
		log.Printf("Сообщение о стиралке отправлено %s (chatID: %d)", userName, chatID)
	}

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		[]tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardButtonData("Забронировать", "order"),
			tgbotapi.NewInlineKeyboardButtonData("Назад", "back"),
		},
	)

	msg.ReplyMarkup = inlineKeyboard

	_, err := bot.Send(msg)
	if err != nil {
		return fmt.Errorf("%w:%w\n", err, errs.ErrSendMessage)
	}

	return nil
}
