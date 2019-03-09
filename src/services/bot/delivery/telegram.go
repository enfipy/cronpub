package delivery

import (
	"errors"
	"strconv"
	"strings"

	"github.com/enfipy/cronpub/src/helpers"
	"github.com/enfipy/cronpub/src/models"
	"github.com/google/uuid"

	telebot "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (server *BotServer) SetupTelegram() {
	updateConfig := telebot.NewUpdate(0)
	updateConfig.Timeout = 60

	updatesChan, err := server.BotInstance.GetUpdatesChan(updateConfig)
	helpers.PanicOnError(err)

	for update := range updatesChan {
		if update.Message == nil {
			continue
		}

		switch {
		case update.Message.Photo != nil:
			server.handle(update.Message, server.SaveImage)
		case update.Message.Video != nil:
			server.handle(update.Message, server.SaveVideo)
		case update.Message.Document != nil:
			server.handle(update.Message, server.SaveGif)
		case strings.Contains(update.Message.Text, "rm "):
			server.handle(update.Message, server.Remove)
		case strings.Contains(update.Message.Text, "send"):
			server.handle(update.Message, server.Send)
		case strings.Contains(update.Message.Text, "fetch"):
			server.handle(update.Message, server.Fetch)
		case strings.Contains(update.Message.Text, "count"):
			server.handle(update.Message, server.Count)
		}
	}
}

func (server *BotServer) SaveImage(msg *telebot.Message) string {
	photos := *msg.Photo
	post := &models.Post{
		FileType:       models.FileType_IMAGE,
		TelegramFileID: photos[0].FileID,
	}
	id := server.BotController.SavePost(post)
	return id.String()
}

func (server *BotServer) SaveVideo(msg *telebot.Message) string {
	post := &models.Post{
		FileType:       models.FileType_VIDEO,
		TelegramFileID: msg.Video.FileID,
	}
	id := server.BotController.SavePost(post)
	return id.String()
}

func (server *BotServer) SaveGif(msg *telebot.Message) string {
	if msg.Document.MimeType != "video/mp4" {
		panic(errors.New("invalid format"))
	}
	post := &models.Post{
		FileType:       models.FileType_GIF,
		TelegramFileID: msg.Document.FileID,
	}
	id := server.BotController.SavePost(post)
	return id.String()
}

func (server *BotServer) Send(_ *telebot.Message) string {
	randomPost := server.BotController.GetRandomPost()
	if randomPost == nil {
		panic(errors.New("no posts"))
	}
	server.sendPost(randomPost)
	return "post sent"
}

func (server *BotServer) Fetch(_ *telebot.Message) string {
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

	return "Success"
}

func (server *BotServer) Count(_ *telebot.Message) string {
	count := server.BotController.CountPosts()
	result := strconv.Itoa(int(count)) + " posts"
	return result
}

func (server *BotServer) Remove(msg *telebot.Message) string {
	del := strings.Split(msg.Text, "rm ")
	if len(del) < 2 {
		return ""
	}

	id, err := uuid.Parse(del[1])
	helpers.PanicOnError(err)

	isDeleted := server.BotController.RemovePost(id)

	var res string
	if isDeleted {
		res = "Post removed"
	} else {
		res = "Post not found"
	}

	return res
}
