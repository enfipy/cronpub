package delivery

import "github.com/tucnak/telebot"

func (server *BotServer) SetupTelegram() {
	server.BotInstance.Handle("hello", server.World)
}

func (server *BotServer) World(m *telebot.Message) {
	server.BotInstance.Reply(m, "world")
}
