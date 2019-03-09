package helpers

import telebot "github.com/go-telegram-bot-api/telegram-bot-api"

func InitTelegram(token string) *telebot.BotAPI {
	botInstance, err := telebot.NewBotAPI(token)
	PanicOnError(err)
	botInstance.Debug = false
	return botInstance
}
