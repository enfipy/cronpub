package bot

import "github.com/enfipy/cronpub/src/models"

type Usecase interface {
	SavePost(post *models.Post)
	GetRandomPost() *models.Post
}
