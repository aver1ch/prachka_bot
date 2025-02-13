package send

import (
	"fmt"
	"laundryBot/internal/errs"
	"laundryBot/internal/processing"
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
			tgbotapi.NewInlineKeyboardButtonData("Сушка", "Сушка"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Стиралка 1", "Стиралка 1"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Стиралка 2", "Стиралка 2"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Стиралка 3", "Стиралка 3"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Мои бронирования", "myOrders"),
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

func SendInfoByService(chatID int64, userName string, bot *tgbotapi.BotAPI, service string) error {
	status, err := processing.ReadInfoFromJSON(service)
	if err != nil {
		return fmt.Errorf("ошибка чтения данных о сервисе для %s: %w", service, err)
	}

	log.Print(status, service)

	var info string
	var inlineKeyboard tgbotapi.InlineKeyboardMarkup

	// Проверка статуса сервиса
	if status.Status == false {
		info = fmt.Sprintf("%s\n\nСтатус: %s\n\nБудет доступно в %02d:%02d:%02d",
			service, "Занято", status.TimeUntilRelease.Hours, status.TimeUntilRelease.Minutes, status.TimeUntilRelease.Seconds)
	} else {
		info = fmt.Sprintf("%s\n\n%s", service, "Свободно")
	}

	msg := tgbotapi.NewMessage(chatID, info)
	log.Printf("Сообщение по сервису %s отправлено %s (chatID: %d)", service, userName, chatID)

	if service == "Сушка" {
		inlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(
			[]tgbotapi.InlineKeyboardButton{
				tgbotapi.NewInlineKeyboardButtonData("Забронировать", "orderDry"),
				tgbotapi.NewInlineKeyboardButtonData("Назад", "back"),
			},
		)
	} else {
		str := fmt.Sprintf("orderLaundry+%s", service) // костыль
		inlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(
			[]tgbotapi.InlineKeyboardButton{
				tgbotapi.NewInlineKeyboardButtonData("Забронировать", str),
				tgbotapi.NewInlineKeyboardButtonData("Назад", "back"),
			},
		)
	}

	msg.ReplyMarkup = inlineKeyboard

	_, err = bot.Send(msg)
	if err != nil {
		return fmt.Errorf("ошибка при отправке сообщения: %w", err)
	}

	return nil
}

func SendRequestOfLaundryMode(chatID int64, userName string, bot *tgbotapi.BotAPI, mode string) error {
	msg := tgbotapi.NewMessage(chatID, "Какой режим будете использовать?\n")

	str1 := fmt.Sprintf("%s+%s", mode, "quick") // костыль
	str2 := fmt.Sprintf("%s+%s", mode, "long")  // костыль

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		[]tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardButtonData("Быстрая стирка", str1),
			tgbotapi.NewInlineKeyboardButtonData("Другой режим", str2),
		},
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Назад", "back"),
		),
	)

	msg.ReplyMarkup = inlineKeyboard

	_, err := bot.Send(msg)
	if err != nil {
		return fmt.Errorf("%w:%w\n", err, errs.ErrSendMessage)
	}

	return nil
}

func SendRequestOfOrderConfirmation(chatID int64, userName string, bot *tgbotapi.BotAPI, service string) error {
	status, err := processing.ReadInfoFromJSON(service)
	if err != nil {
		return fmt.Errorf("%w:%w\n", err, errs.ErrReadStatusFile)
	}

	info := fmt.Sprintf("Вы уверены, что хотите совершить бронь на %02d:%02d:%02d?",
		status.TimeUntilRelease.Hours, status.TimeUntilRelease.Minutes, status.TimeUntilRelease.Seconds)

	msg := tgbotapi.NewMessage(chatID, info)
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		[]tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardButtonData("Подтвердить", "сonfirm"),
			tgbotapi.NewInlineKeyboardButtonData("Назад", "back"),
		},
	)

	msg.ReplyMarkup = inlineKeyboard

	_, err = bot.Send(msg)
	if err != nil {
		return fmt.Errorf("%w:%w\n", err, errs.ErrSendMessage)
	}

	return nil
}

func SendConfirmMessage(chatID int64, userName string, bot *tgbotapi.BotAPI) error {
	info := fmt.Sprintf("Бронь подтверждена, если машинка не занята, то ты можешь начинать стирку, в противном случае жди от меня сообщения, когда твоя машинка освободится\nУ тебя будет 5 минут, чтобы начать стирку")

	msg := tgbotapi.NewMessage(chatID, info)
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Хорошо", "back"),
		),
	)

	msg.ReplyMarkup = inlineKeyboard

	_, err := bot.Send(msg)
	if err != nil {
		return fmt.Errorf("%w:%w\n", err, errs.ErrSendMessage)
	}

	return nil
}

func SendMyOrders(chatID int64, userName string, bot *tgbotapi.BotAPI) error {
	info := fmt.Sprintf("Твои бронирования: \n\n Выбери, на какую из машинок ты хочешь отменить свою бронь?\n")

	msg := tgbotapi.NewMessage(chatID, info)
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Сушка", "XСушка"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Стиралка 1", "XСтиралка 1"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Стиралка 2", "XСтиралка 2"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Стиралка 3", "XСтиралка 3"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Назад", "back"),
		),
	)

	msg.ReplyMarkup = inlineKeyboard

	_, err := bot.Send(msg)
	if err != nil {
		return fmt.Errorf("%w:%w\n", err, errs.ErrSendMessage)
	}

	return nil
}
