package bot

import "github.com/enfipy/cronpub/src/models"

type Controller interface {
	SavePost(post *models.Post)
	GetRandomPost() *models.Post
}
