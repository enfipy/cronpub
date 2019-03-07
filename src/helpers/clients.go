package helpers

import (
	"time"

	"github.com/tucnak/telebot"
)

func InitTelegram(token string) *telebot.Bot {
	poller := &telebot.LongPoller{
		Timeout: 0 * time.Second,
	}
	settings := telebot.Settings{
		Token:  token,
		Poller: poller,
	}
	botInstance, err := telebot.NewBot(settings)
	PanicOnError(err)
	return botInstance
}
