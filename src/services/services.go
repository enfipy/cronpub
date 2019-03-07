package services

import (
	"errors"
	"log"

	"github.com/enfipy/cronpub/src/config"
	"github.com/enfipy/cronpub/src/helpers"

	cron "github.com/robfig/cron"
)

func InitServices(cnfg *config.Config) (start, close func()) {
	if cnfg.Settings == nil {
		helpers.PanicOnError(errors.New("Valid settings must be provided"))
	}

	pool := helpers.InitRedis(cnfg.RedisAddress, cnfg.RedisNetwork)
	botInstance := helpers.InitTelegram(cnfg.Settings.Telegram.BotToken)
	cronInstance := cron.New()

	log.Print(pool, botInstance, cronInstance)

	start = func() {}
	close = func() {}
	return
}
