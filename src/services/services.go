package services

import (
	"errors"
	"net/http"
	"time"

	"github.com/enfipy/cronpub/src/config"
	"github.com/enfipy/cronpub/src/helpers"

	botController "github.com/enfipy/cronpub/src/services/bot/controller"
	botDelivery "github.com/enfipy/cronpub/src/services/bot/delivery"
	botUsecase "github.com/enfipy/cronpub/src/services/bot/usecase"

	scraperController "github.com/enfipy/cronpub/src/services/scraper/controller"
	scraperDelivery "github.com/enfipy/cronpub/src/services/scraper/delivery"

	"github.com/enfipy/locker"
	"github.com/robfig/cron"
)

func InitServices(cnfg *config.Config) (start, close func()) {
	if cnfg.Settings == nil {
		helpers.PanicOnError(errors.New("Valid settings must be provided"))
	}

	locker := locker.Initialize()
	pool := helpers.InitRedis(cnfg.RedisAddress, cnfg.RedisNetwork)
	botInstance := helpers.InitTelegram(cnfg.Settings.Telegram.BotToken)
	cronInstance := cron.New()
	httpClient := &http.Client{
		Timeout: time.Second * 5,
	}

	ucsBot := botUsecase.NewUsecase(pool, locker)

	cnrBot := botController.NewController(ucsBot)
	cnrScraper := scraperController.NewController(httpClient)

	scraperDelivery.NewDelivery(cnfg, cnrScraper, cnrBot, cronInstance)
	dlrBot := botDelivery.NewDelivery(cnfg, cnrBot, cnrScraper, botInstance, cronInstance)

	start = func() {
		dlrBot.SetupCron()
		cronInstance.Start()
		dlrBot.SetupTelegram()
	}
	close = func() {
		cronInstance.Stop()
		botInstance.StopReceivingUpdates()
	}
	return
}
