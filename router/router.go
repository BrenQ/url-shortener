package router

import (
	"github.com/gin-gonic/gin"
	"urlshortener/controller"
)

// Get a router with endpoints for shorten url

type Router struct {
	LinkController controller.LinkControllerInterface
}

// Return a new router instance
func New () Router {
	return Router {
			LinkController: controller.NewLinkController(),
	}
}

func ( router * Router) Get() * gin.Engine  {
	r := gin.Default()
	// Group routes with prefix
	//v := r.Group("/api/v1")
	// Create a shortened url
	r.POST("/links" , router.LinkController.CreateLink)
	// Get by shortened url and redirect to the original
	r.GET("/:code" , router.LinkController.GetLink)
	return r
}
