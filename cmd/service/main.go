/*

функционал таков: сразу будет вопрос стираться или сушиться
в ответ на кнопку будет информация о бронях и кнопка забронировать на такое-то время
какая очередь, количество людей, на сколько циклов заняли, какой режим стирки, как раз от этого и зависит время
Можно отменить бронь
Напоминалка по истечению таймера
*/

package main

import (
	"laundryBot/internal/handlers"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	bot, err := tgbotapi.NewBotAPI("7633580945:AAGSE71kwZY7zNLn6HpO-eNDCj_ViX1pCBk")
	if err != nil {
		log.Panicf("Ошибка при запуске бота: %v", err)
	}

	log.Printf("Бот запущен\nАвторизован как %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {

			err := handlers.HandleStartButton(update, bot)
			if err != nil {
				log.Printf("Ошибка в обработке команды 'start': %v", err)
			}

			/* err = handlers.HandleVerifyButton(update.CallbackQuery, bot)
			if err != nil {
				log.Printf("Ошибка в обработке номера комнаты: %v", err)
			} */
		}

		if update.CallbackQuery != nil {

			err := handlers.HandleCallbackQuery(update.CallbackQuery, bot)
			if err != nil {
				log.Printf("Ошибка в обработке CallbackQuery: %v", err)
			}

		}
	}
}
