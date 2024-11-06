package main

import (
	"log"

	tgconfighelper "github.com/BlinovDev/go-tg-bot-config-helper"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	// Замените "YOUR_TELEGRAM_BOT_TOKEN" на токен, который вы получили от BotFather
	config, err := tgconfighelper.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	bot, err := tgbotapi.NewBotAPI(config.Bot.Token)
	if err != nil {
		log.Panic(err)
	}

	// Включаем логирование для отладки
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Получаем обновления от API
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// Обработка сообщений
	for update := range updates {
		if update.Message != nil { // Проверяем, что это текстовое сообщение
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			// Ответ на сообщение пользователя
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			if update.Message.Text == "/start" {
				msg.Text = "Добро пожаловать! Я ваш Telegram бот на Go."
			}

			bot.Send(msg)
		}
	}
}
