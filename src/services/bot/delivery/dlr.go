package delivery

import (
	"github.com/enfipy/cronpub/src/config"
	"github.com/enfipy/cronpub/src/helpers"
	"github.com/enfipy/cronpub/src/services/bot"
	"github.com/enfipy/cronpub/src/services/scraper"

	telebot "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/robfig/cron"
)

type BotServer struct {
	Config            *config.Config
	BotController     bot.Controller
	ScraperController scraper.Controller

	BotInstance  *telebot.BotAPI
	CronInstance *cron.Cron
}

func NewDelivery(cnfg *config.Config, cnrBot bot.Controller, cnrScraper scraper.Controller, botInstance *telebot.BotAPI, cronInstance *cron.Cron) *BotServer {
	server := &BotServer{
		Config:            cnfg,
		BotController:     cnrBot,
		ScraperController: cnrScraper,

		BotInstance:  botInstance,
		CronInstance: cronInstance,
	}

	return server
}

func (server *BotServer) SetupCron() {
	for _, cron := range server.Config.Settings.Telegram.Crons {
		server.CronInstance.AddFunc(cron, func() {
			defer helpers.RecoverWithLog()
			server.Send(nil)
		})
	}
}
