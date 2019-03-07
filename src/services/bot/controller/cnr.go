package controller

import (
	"github.com/enfipy/cronpub/src/services/bot"
)

type botController struct {
	botUsecase bot.Usecase
}

func NewController(botUsecase bot.Usecase) bot.Controller {
	return &botController{
		botUsecase: botUsecase,
	}
}
