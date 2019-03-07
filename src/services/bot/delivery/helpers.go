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

	switch post.FileType {
	case models.FileType_GIF:
		file, err := server.BotInstance.FileByID(post.TelegramFileID)
		helpers.PanicOnError(err)

		sendable = &telebot.Document{
			File:    file,
			Caption: "@epiocus",
		}
	case models.FileType_VIDEO:
		file, err := server.BotInstance.FileByID(post.TelegramFileID)
		helpers.PanicOnError(err)

		sendable = &telebot.Video{
			File:    file,
			Caption: "@epiocus",
		}
	case models.FileType_IMAGE:
		file := telebot.FromURL(post.FileLink)
		sendable = &telebot.Photo{
			File:    file,
			Caption: "@epiocus",
		}
	}

	return sendable
}
