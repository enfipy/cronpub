package delivery

import (
	"github.com/enfipy/cronpub/src/helpers"
	"github.com/enfipy/cronpub/src/models"

	"github.com/tucnak/telebot"
)

func (server *BotServer) handle(logic func(msg *telebot.Message)) func(msg *telebot.Message) {
	return func(msg *telebot.Message) {
		defer func() {
			if rec := recover(); rec != nil {
				err := rec.(error)
				errorMessage := "Error: " + err.Error()
				server.BotInstance.Reply(msg, errorMessage)
			}
		}()
		// Todo: Authenticate user
		logic(msg)
	}
}

func (server *BotServer) getSendable(post *models.Post) telebot.Sendable {
	var sendable telebot.Sendable

	if post.FileLink != "" {
		sendable = &telebot.Photo{
			File:    telebot.FromURL(post.FileLink),
			Caption: server.Config.Settings.Telegram.Caption,
		}
		return sendable
	}

	file, err := server.BotInstance.FileByID(post.TelegramFileID)
	helpers.PanicOnError(err)

	switch post.FileType {
	case models.FileType_GIF:
		sendable = &telebot.Document{
			File:    file,
			Caption: server.Config.Settings.Telegram.Caption,
		}
	case models.FileType_VIDEO:
		sendable = &telebot.Video{
			File:    file,
			Caption: server.Config.Settings.Telegram.Caption,
		}
	case models.FileType_IMAGE:
		sendable = &telebot.Photo{
			File:    file,
			Caption: server.Config.Settings.Telegram.Caption,
		}
	}

	return sendable
}

func (server *BotServer) sendPost(post *models.Post) {
	sendable := server.getSendable(post)
	chat := &telebot.Chat{
		Username: server.Config.Settings.Telegram.ChatName,
		Type:     telebot.ChatChannel,
	}
	server.BotInstance.Send(chat, sendable)
}
