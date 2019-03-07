package config

type CronpubSettings struct {
	Telegram struct {
		BotToken string `yaml:"bot_token"`
		Caption  string `yaml:"caption"`
		ChatName string `yaml:"chat_name"`
	} `yaml:"telegram"`
}
