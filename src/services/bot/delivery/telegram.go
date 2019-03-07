package delivery

import (
	"errors"
	"strconv"
	"strings"

	"github.com/enfipy/cronpub/src/helpers"

	"github.com/enfipy/cronpub/src/models"
	"github.com/google/uuid"

	"github.com/tucnak/telebot"
)

func (server *BotServer) SetupTelegram() {
	server.BotInstance.Handle(telebot.OnPhoto, server.handle(server.SaveImage))
	server.BotInstance.Handle(telebot.OnVideo, server.handle(server.SaveVideo))
	server.BotInstance.Handle(telebot.OnDocument, server.handle(server.SaveGif))
	server.BotInstance.Handle("send", server.handle(server.Send))
	server.BotInstance.Handle("fetch", server.handle(server.Fetch))
	server.BotInstance.Handle("count", server.handle(server.Count))
	server.BotInstance.Handle(telebot.OnText, server.handle(server.OnText))
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
	server.sendPost(randomPost)
}

func (server *BotServer) Fetch(_ *telebot.Message) {
	scraperLinks := server.Config.Settings.Scraper.Links
	if len(scraperLinks) <= 0 {
		panic(errors.New("no links"))
	}

	for _, link := range scraperLinks {
		links := server.ScraperController.FetchFromReddit(link)

		for _, fileLink := range links {
			post := &models.Post{
				FileType: models.FileType_IMAGE,
				FileLink: fileLink,
			}
			server.BotController.SavePost(post)
		}
	}
}

func (server *BotServer) Count(msg *telebot.Message) {
	count := server.BotController.CountPosts()
	result := strconv.Itoa(int(count)) + " posts"
	server.BotInstance.Reply(msg, result)
}

func (server *BotServer) OnText(msg *telebot.Message) {
	del := strings.Split(msg.Text, "rm ")
	if len(del) >= 2 {
		server.Remove(msg, del[1])
	}
}

func (server *BotServer) Remove(msg *telebot.Message, delID string) {
	id, err := uuid.Parse(delID)
	helpers.PanicOnError(err)

	isDeleted := server.BotController.RemovePost(id)

	var res string
	if isDeleted {
		res = "Post removed"
	} else {
		res = "Post not found"
	}

	server.BotInstance.Reply(msg, res)
}
