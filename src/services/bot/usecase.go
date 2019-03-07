package bot

import (
	"github.com/enfipy/cronpub/src/models"

	"github.com/google/uuid"
)

type Usecase interface {
	SavePost(post *models.Post) uuid.UUID
	GetRandomPost() *models.Post
	CountPosts() int64
	RemovePost(id uuid.UUID) bool
}
