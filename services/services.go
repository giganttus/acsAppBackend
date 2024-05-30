package services

import (
	"acsApp/models"
	"acsApp/repos"

	"github.com/gin-gonic/gin"
)

type Services struct {
	Repos repos.IRepos
}

func NewServices(repos repos.IRepos) *Services {
	return &Services{
		Repos: repos,
	}
}

// Function located in services, used by handlers
type IServices interface {
	// Posts
	GetPostById(*gin.Context) (models.Posts, error)
	GetFeaturedPosts(*gin.Context) ([]models.PostsPreview, error)
	GetPostsPerPage(*gin.Context) ([]models.PostsPreview, error)
}
