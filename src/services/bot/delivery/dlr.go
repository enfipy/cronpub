package delivery

import (
	"github.com/enfipy/cronpub/src/config"

	"github.com/enfipy/cronpub/src/services/bot"

	"github.com/robfig/cron"
	"github.com/tucnak/telebot"
)

type BotServer struct {
	Config        *config.Config
	BotController bot.Controller

	BotInstance  *telebot.Bot
	CronInstance *cron.Cron
}

func NewDelivery(cnfg *config.Config, cnrBot bot.Controller, botInstance *telebot.Bot, cronInstance *cron.Cron) *BotServer {
	server := &BotServer{
		Config:        cnfg,
		BotController: cnrBot,

		BotInstance:  botInstance,
		CronInstance: cronInstance,
	}

	server.SetupTelegram()
	server.SetupCron()

	return server
}

func (server *BotServer) SetupCron() {
	// server.CronInstance.AddFunc("*", server.BotController.World)
	// server.CronInstance.AddFunc("*", server.BotController.World)
	// server.CronInstance.AddFunc("*", server.BotController.World)
}
