package delivery

import (
	"github.com/enfipy/cronpub/src/helpers"
	"github.com/enfipy/cronpub/src/models"

	telebot "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (server *BotServer) handle(msg *telebot.Message, handler func(msg *telebot.Message) string) {
	defer func() {
		if rec := recover(); rec != nil {
			err := rec.(error)
			errorMessage := "Error: " + err.Error()

			res := telebot.NewMessage(msg.Chat.ID, errorMessage)
			res.ReplyToMessageID = msg.MessageID

			server.BotInstance.Send(res)
		}
	}()
	// Todo: Authenticate user
	result := handler(msg)

	if result != "" {
		res := telebot.NewMessage(msg.Chat.ID, result)
		res.ReplyToMessageID = msg.MessageID

		server.BotInstance.Send(res)
	}
}

func (server *BotServer) getChat() *telebot.Chat {
	chatConfig := telebot.ChatConfig{
		SuperGroupUsername: server.Config.Settings.Telegram.ChatName,
	}
	chat, err := server.BotInstance.GetChat(chatConfig)
	helpers.PanicOnError(err)
	return &chat
}

func (server *BotServer) sendPost(post *models.Post) {
	var msg telebot.Chattable
	chat := server.getChat()

	if post.FileLink != "" {
		media := telebot.NewInputMediaPhoto(post.FileLink)
		media.Caption = server.Config.Settings.Telegram.Caption
		msg = telebot.NewMediaGroup(chat.ID, []interface{}{media})
	} else {
		switch post.FileType {
		case models.FileType_GIF:
			doc := telebot.NewDocumentShare(chat.ID, post.TelegramFileID)
			doc.Caption = server.Config.Settings.Telegram.Caption
			msg = doc
		case models.FileType_VIDEO:
			video := telebot.NewVideoShare(chat.ID, post.TelegramFileID)
			video.Caption = server.Config.Settings.Telegram.Caption
			msg = video
		case models.FileType_IMAGE:
			image := telebot.NewPhotoShare(chat.ID, post.TelegramFileID)
			image.Caption = server.Config.Settings.Telegram.Caption
			msg = image
		}
	}

	_, err := server.BotInstance.Send(msg)
	helpers.PanicOnError(err)
}
