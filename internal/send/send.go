package send

import (
	"fmt"
	"laundryBot/internal/errs"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendStartMessage(chatID int64, userName string, bot *tgbotapi.BotAPI) error {
	msg := tgbotapi.NewMessage(chatID, "Привет! Этот бот поможет решить проблему со стиралками, теперь тебе не нужно постоянно переписываться в чатике по поводу стиралок.\nДостаточно нажать одну кнопку в этом боте, чтобы занять стиралку или сушилку :)")
	_, err := bot.Send(msg)
	if err != nil {
		return fmt.Errorf("%w:%w\n", err, errs.ErrSendMessage)
	}
	log.Printf("The start message sended to %s (chatID: %d)", userName, chatID)
	return nil
}
