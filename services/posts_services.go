package services

import (
	"acsApp/helpers"
	"acsApp/models"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Services) GetPostById(ctx *gin.Context) (models.Posts, error) {
	id := ctx.Query("id")
	postId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("postid", postId)
		return models.Posts{}, err
	}

	return s.Repos.GetPostById(postId)
}

func (s *Services) GetFeaturedPosts(ctx *gin.Context) ([]models.PostsPreview, error) {
	sliNum := os.Getenv("SLI_NUM")
	if sliNum == "" {
		return []models.PostsPreview{}, helpers.ErrEnvLoad
	}

	sliNumConv, err := strconv.Atoi(sliNum)
	if err != nil {
		return []models.PostsPreview{}, helpers.ErrEnvConvert
	}

	return s.Repos.GetFeaturedPosts(sliNumConv)
}

func (s *Services) GetPostsPerPage(ctx *gin.Context) ([]models.PostsPreview, error) {
	id := ctx.Query("id")
	pageId, err := strconv.Atoi(id)
	if err != nil {
		return []models.PostsPreview{}, err
	}

	ppNum := os.Getenv("PP_NUM")
	if ppNum == "" {
		return []models.PostsPreview{}, helpers.ErrEnvLoad
	}

	ppNumConv, err := strconv.Atoi(ppNum)
	if err != nil {
		return []models.PostsPreview{}, helpers.ErrEnvConvert
	}

	return s.Repos.GetPostsPerPage(pageId, ppNumConv)
}
