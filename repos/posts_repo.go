package repos

import (
	"acsApp/helpers"
	"acsApp/models"
)

// Main functions
func (r *Repos) GetPostById(postId int) (models.Posts, error) {
	var post models.Posts

	if err := r.db.First(&post, postId).Error; err != nil {
		return post, helpers.ErrRepo
	}

	return post, nil
}

func (r *Repos) GetFeaturedPosts(sliNum int) ([]models.PostsPreview, error) {
	var postsPreview []models.PostsPreview

	if res := r.db.Model(&models.Posts{}).Select("id, author, category, date, featured, ilink, subtitle, tags, title").Where("featured = 1").Order("id desc").Limit(sliNum).Scan(&postsPreview); res.Error != nil {
		return nil, helpers.ErrRepo
	}

	return postsPreview, nil
}

func (r *Repos) GetPostsPerPage(pageId int, ppNum int) ([]models.PostsPreview, error) {
	var postsPreview []models.PostsPreview

	if res := r.db.Model(&models.Posts{}).Select("id, author, category, date, featured, ilink, subtitle, tags, title").Order("id desc").Limit(ppNum).Offset((pageId - 1) * ppNum).Scan(&postsPreview); res.Error != nil {
		return nil, helpers.ErrRepo
	}

	return postsPreview, nil
}
