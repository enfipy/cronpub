package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/enfipy/cronpub/src/config"
	"github.com/enfipy/cronpub/src/helpers"
	"github.com/enfipy/cronpub/src/services"
)

func main() {
	cnfg := config.InitConfig()
	cnfg.Settings = helpers.GetSettings("/settings.yaml")

	if cnfg.AppEnv != "production" {
		log.SetFlags(0)
	} else {
		log.Print("Started")
	}

	start, close := services.InitServices(cnfg)

	go gracefulShutdown(close)
	start()
}

func gracefulShutdown(close func()) {
	quitChan := make(chan os.Signal, 1)

	signal.Notify(quitChan, syscall.SIGTERM)
	signal.Notify(quitChan, syscall.SIGINT)
	signal.Notify(quitChan, syscall.SIGKILL)

	<-quitChan
	close()
}
