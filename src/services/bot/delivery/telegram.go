package delivery

import (
	"errors"

	"github.com/enfipy/cronpub/src/models"

	"github.com/tucnak/telebot"
)

func (server *BotServer) SetupTelegram() {
	server.BotInstance.Handle(telebot.OnVideo, server.handle(server.SaveVideo))
	server.BotInstance.Handle(telebot.OnDocument, server.handle(server.SaveGif))
	server.BotInstance.Handle("get", server.handle(server.Get))
}

func (server *BotServer) SaveVideo(msg *telebot.Message) {
	post := &models.Post{
		FileType:       models.FileType_VIDEO,
		TelegramFileID: msg.Video.FileID,
	}
	server.BotController.SavePost(post)
	server.BotInstance.Reply(msg, "Video saved")
}

func (server *BotServer) SaveGif(msg *telebot.Message) {
	if msg.Document.MIME != "video/mp4" {
		panic(errors.New("invalid format"))
	}
	post := &models.Post{
		FileType:       models.FileType_GIF,
		TelegramFileID: msg.Document.FileID,
	}
	server.BotController.SavePost(post)
	server.BotInstance.Reply(msg, "Gif saved")
}

func (server *BotServer) Get(_ *telebot.Message) {
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
