package delivery

import (
	"errors"

	"github.com/enfipy/cronpub/src/models"

	"github.com/tucnak/telebot"
)

func (server *BotServer) SetupTelegram() {
	server.BotInstance.Handle(telebot.OnPhoto, server.handle(server.SaveImage))
	server.BotInstance.Handle(telebot.OnVideo, server.handle(server.SaveVideo))
	server.BotInstance.Handle(telebot.OnDocument, server.handle(server.SaveGif))
	server.BotInstance.Handle("send", server.handle(server.Send))
}

func (server *BotServer) SaveImage(msg *telebot.Message) {
	post := &models.Post{
		FileType:       models.FileType_IMAGE,
		TelegramFileID: msg.Photo.FileID,
	}
	id := server.BotController.SavePost(post)
	server.BotInstance.Reply(msg, id.String())
}

func (server *BotServer) SaveVideo(msg *telebot.Message) {
	post := &models.Post{
		FileType:       models.FileType_VIDEO,
		TelegramFileID: msg.Video.FileID,
	}
	id := server.BotController.SavePost(post)
	server.BotInstance.Reply(msg, id.String())
}

func (server *BotServer) SaveGif(msg *telebot.Message) {
	if msg.Document.MIME != "video/mp4" {
		panic(errors.New("invalid format"))
	}
	post := &models.Post{
		FileType:       models.FileType_GIF,
		TelegramFileID: msg.Document.FileID,
	}
	id := server.BotController.SavePost(post)
	server.BotInstance.Reply(msg, id.String())
}

func (server *BotServer) Send(_ *telebot.Message) {
	randomPost := server.BotController.GetRandomPost()
	if randomPost == nil {
		panic(errors.New("no posts"))
	}
	server.SendPost(randomPost)
}

func (server *BotServer) SendPost(post *models.Post) {
	sendable := server.getSendable(post)
	chat := &telebot.Chat{
		Username: server.Config.Settings.Telegram.ChatName,
		Type:     telebot.ChatChannel,
	}
	server.BotInstance.Send(chat, sendable)
}
