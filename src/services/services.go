package services

import (
	"errors"

	"github.com/enfipy/cronpub/src/config"
	"github.com/enfipy/cronpub/src/helpers"

	botController "github.com/enfipy/cronpub/src/services/bot/controller"
	botDelivery "github.com/enfipy/cronpub/src/services/bot/delivery"
	botUsecase "github.com/enfipy/cronpub/src/services/bot/usecase"

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

	ucsBot := botUsecase.NewUsecase(pool, locker)
	cnrBot := botController.NewController(ucsBot)
	botDelivery.NewDelivery(cnfg, cnrBot, botInstance, cronInstance)

	start = func() {
		cronInstance.Start()
		botInstance.Start()
	}
	close = func() {
		cronInstance.Stop()
		botInstance.Stop()
	}
	return
}
