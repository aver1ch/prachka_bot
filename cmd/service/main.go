package main

import (
	"laundryBot/internal/handlers"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

/*
нужна верификация
Номер комнаты, id в тг, проверка на подписку в беседу в общаге
закидываем это всё дело в базу данных, если всё хорошо
после этого будет доступен функционал:
функционал таков: Сразу будет сообщение, которое будет изменяться динамически или по кнопке
там будет показано, какая машинка, кем занята или не занята, поломана ли машинка
какая очередь, количество людей, на сколько циклов заняли, какой режим стирки, как раз от этого и зависит время
Можно будет занять наперёд, если не мешает другой стирке
Можно отменить бронь
Напоминалка по истечению таймера
*/

func main() {
	bot, err := tgbotapi.NewBotAPI("7633580945:AAGSE71kwZY7zNLn6HpO-eNDCj_ViX1pCBk")

	if err != nil {
		log.Panic("Bot is not started: %w\n", err)
	}

	bot.Debug = true

	log.Printf("Bot started\nAutorised as %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates := bot.GetUpdatesChan(u)

	err = handlers.HandleStartButton(updates, bot)
	if err != nil {
		log.Panic(err)
	}

	err = handlers.HandleVerificationButton(updates, bot)
	if err != nil {
		log.Panic(err)
	}
}
