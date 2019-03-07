package controller

import (
	"github.com/enfipy/cronpub/src/models"
	"github.com/enfipy/cronpub/src/services/bot"

	"github.com/google/uuid"
)

type botController struct {
	botUsecase bot.Usecase
}

func NewController(botUsecase bot.Usecase) bot.Controller {
	return &botController{
		botUsecase: botUsecase,
	}
}

func (cnr *botController) SavePost(post *models.Post) uuid.UUID {
	return cnr.botUsecase.SavePost(post)
}

func (cnr *botController) GetRandomPost() *models.Post {
	return cnr.botUsecase.GetRandomPost()
}
