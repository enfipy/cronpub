package delivery

import (
	"errors"

	"github.com/enfipy/cronpub/src/config"
	"github.com/enfipy/cronpub/src/helpers"
	"github.com/enfipy/cronpub/src/models"
	"github.com/enfipy/cronpub/src/services/bot"
	"github.com/enfipy/cronpub/src/services/scraper"

	"github.com/robfig/cron"
)

type ScraperServer struct {
	Config            *config.Config
	ScraperController scraper.Controller
	BotController     bot.Controller
	CronInstance      *cron.Cron
}

func NewDelivery(cnfg *config.Config, cnrScraper scraper.Controller, cnrBot bot.Controller, cronInstance *cron.Cron) *ScraperServer {
	server := &ScraperServer{
		Config:            cnfg,
		ScraperController: cnrScraper,
		BotController:     cnrBot,
		CronInstance:      cronInstance,
	}

	server.SetupCron()

	return server
}

func (server *ScraperServer) SetupCron() {
	for _, cron := range server.Config.Settings.Scraper.Crons {
		server.CronInstance.AddFunc(cron, func() {
			defer helpers.RecoverWithLog()
			server.Fetch()
		})
	}
}

func (server *ScraperServer) Fetch() {
	// Todo: Combine with bot service fetch func

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
