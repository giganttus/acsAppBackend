package handlers

import (
	"acsApp/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) GetPostById(ctx *gin.Context) {
	res, err := h.Services.GetPostById(ctx)
	if err != nil {
		if err.Error() == helpers.ErrForbidden.Error() {
			ctx.IndentedJSON(http.StatusForbidden, gin.H{"response": "error occured"})
			return
		}

		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"response": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"response": res})
}

func (h *Handlers) GetFeaturedPosts(ctx *gin.Context) {
	res, err := h.Services.GetFeaturedPosts(ctx)
	if err != nil {
		if err.Error() == helpers.ErrForbidden.Error() {
			ctx.IndentedJSON(http.StatusForbidden, gin.H{"response": "error occured"})
			return
		}

		ctx.IndentedJSON(http.StatusForbidden, gin.H{"response": "error occured"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"response": res})
}

func (h *Handlers) GetPostsPerPage(ctx *gin.Context) {
	res, err := h.Services.GetPostsPerPage(ctx)
	if err != nil {
		if err.Error() == helpers.ErrForbidden.Error() {
			ctx.IndentedJSON(http.StatusForbidden, gin.H{"response": "error occured"})
			return
		}

		ctx.IndentedJSON(http.StatusForbidden, gin.H{"response": "error occured"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"response": res})
}
