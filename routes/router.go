package routes

import (
	"acsApp/handlers"
	"acsApp/middlewares"
	"acsApp/repos"
	"acsApp/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(db *gorm.DB, rate int, burst int) *gin.Engine {
	router := gin.Default()
	router.Use(
		middlewares.CorsMiddleware(),
		middlewares.LoggerMiddleware(),
		middlewares.RateLimiterPerIP(rate, burst),
	)

	rep := repos.NewRepos(db)
	svc := services.NewServices(rep)
	hlr := handlers.NewHandlers(svc)

	apiRouter := router.Group("/api")

	// Posts
	/* noAuthRouter.GET("admin", hlr.GetDailyOffers) */ //Fake, trolling
	apiRouter.GET("post", hlr.GetPostById)
	apiRouter.GET("posts/featured", hlr.GetFeaturedPosts)
	apiRouter.GET("posts", hlr.GetPostsPerPage)

	return router
}
