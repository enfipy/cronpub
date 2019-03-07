package controller

import (
	"github.com/enfipy/cronpub/src/models"
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

func (cnr *botController) SavePost(post *models.Post) {
	cnr.botUsecase.SavePost(post)
}

func (cnr *botController) GetRandomPost() *models.Post {
	return cnr.botUsecase.GetRandomPost()
}
