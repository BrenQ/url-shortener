package router

import (
	"github.com/gin-gonic/gin"
	"urlshortener/controller"
)

// Get a router with endpoints for shorten url

func Get() * gin.Engine  {
	r := gin.Default()
	// Group routes with prefix
	v := r.Group("/api/v1")

	// Create a shortened url
	v.POST("/links" , controller.CreateLink)
	// Get by shortened url and redirect to the original
	v.GET("/links/:id" , controller.GetLink)
	return r
}
