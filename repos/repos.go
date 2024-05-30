package repos

import (
	"acsApp/models"

	"gorm.io/gorm"
)

type Repos struct {
	db *gorm.DB
}

func NewRepos(db *gorm.DB) *Repos {
	return &Repos{
		db: db,
	}
}

// Function located in repo, used by services
type IRepos interface {
	// Posts - Main functions
	GetPostById(int) (models.Posts, error)
	GetFeaturedPosts(int) ([]models.PostsPreview, error)
	GetPostsPerPage(int, int) ([]models.PostsPreview, error)
}
