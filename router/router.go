package router

import (
	"github.com/gin-gonic/gin"
	"urlshortener/controller"
)


func Get() * gin.Engine  {
	r := gin.Default()

	v := r.Group("/api/v1")
	v.POST("/links" , controller.CreateLink)
	v.GET("/links/:id" , controller.GetLink)
	return r
}
