package bot

import (
	"github.com/enfipy/cronpub/src/models"

	"github.com/google/uuid"
)

type Controller interface {
	SavePost(post *models.Post) uuid.UUID
	GetRandomPost() *models.Post
}
