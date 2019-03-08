package delivery

import (
	"os"

	"github.com/enfipy/cronpub/src/config"
	"github.com/enfipy/cronpub/src/helpers"
	"github.com/enfipy/cronpub/src/services/bot"
	"github.com/enfipy/cronpub/src/services/scraper"

	"github.com/robfig/cron"
	"github.com/tucnak/telebot"
)

type BotServer struct {
	Config            *config.Config
	BotController     bot.Controller
	ScraperController scraper.Controller

	BotInstance  *telebot.Bot
	CronInstance *cron.Cron
}

func NewDelivery(cnfg *config.Config, cnrBot bot.Controller, cnrScraper scraper.Controller, botInstance *telebot.Bot, cronInstance *cron.Cron) *BotServer {
	server := &BotServer{
		Config:            cnfg,
		BotController:     cnrBot,
		ScraperController: cnrScraper,

		BotInstance:  botInstance,
		CronInstance: cronInstance,
	}

	server.SetupTelegram()
	server.SetupCron()

	return server
}

func (server *BotServer) SetupCron() {
	for _, cron := range server.Config.Settings.Telegram.Crons {
		server.CronInstance.AddFunc(cron, func() {
			defer helpers.RecoverWithLog()
			server.Send(nil)
			// Todo: Fix lib and this line
			os.Exit(1)
		})
	}
}
