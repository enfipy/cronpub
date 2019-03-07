package config

import "os"

type Config struct {
	AppEnv string

	RedisAddress string
	RedisNetwork string

	Settings *CronpubSettings
}

func InitConfig() *Config {
	cnfg := new(Config)

	cnfg.AppEnv = os.Getenv("APP_ENV")
	cnfg.RedisAddress = os.Getenv("REDIS_ADDRESS")
	cnfg.RedisNetwork = os.Getenv("REDIS_NETWORK")

	return cnfg
}
