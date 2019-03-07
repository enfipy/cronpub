package config

type CronpubSettings struct {
	Telegram struct {
		BotToken string   `yaml:"bot_token"`
		Caption  string   `yaml:"caption"`
		ChatName string   `yaml:"chat_name"`
		Crons    []string `yaml:"crons"`
	} `yaml:"telegram"`

	Scraper struct {
		Links []string `yaml:"links"`
		Crons []string `yaml:"crons"`
	} `yaml:"scraper"`
}
