package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var bot *tgbotapi.BotAPI
var chatId int64

func connectWithTelegram() {
	var err error
	if bot, err = tgbotapi.NewBotAPI(TOKEN); err != nil {
		panic("Cannot connect to Telefram")
	}
}

func sendMessage(msg string) {
	msgConfig := tgbotapi.NewMessage(chatId, msg)
	bot.Send(msgConfig)
}
func main() {
	connectWithTelegram()

	updeteConfig := tgbotapi.NewUpdate(0)
	for update := range bot.GetUpdatesChan(updeteConfig) {
		chatId = update.Message.Chat.ID
		if update.Message != nil && update.Message.Text == "/start" {
			sendMessage("тест ответа")
		}
	}
}
