package delivery

import (
	"github.com/enfipy/cronpub/src/config"

	botService "github.com/enfipy/cronpub/src/services/bot"

	"github.com/robfig/cron"
	"github.com/tucnak/telebot"
)

type BotServer struct {
	Config        *config.Config
	BotController botService.Controller

	BotInstance  *telebot.Bot
	CronInstance *cron.Cron
}

func NewDelivery(cnfg *config.Config, cnrBot botService.Controller, botInstance *telebot.Bot, cronInstance *cron.Cron) *BotServer {
	return &BotServer{
		Config:        cnfg,
		BotController: cnrBot,

		BotInstance:  botInstance,
		CronInstance: cronInstance,
	}
}

func (server *BotServer) SetupCron() {
	// server.CronInstance.AddFunc("*", server.BotController.World)
	// server.CronInstance.AddFunc("*", server.BotController.World)
	// server.CronInstance.AddFunc("*", server.BotController.World)
}
