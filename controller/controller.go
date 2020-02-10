package controller

import (
	"github.com/gin-gonic/gin"
	"urlshortener/service"
)

type Controller struct {

}

// Definition LinkController
type LinkControllerInterface interface {
	CreateLink(c * gin.Context)
	GetLink(c *gin.Context)
}

type LinkController struct {
	LinkService service.LinkServiceInterface
}

// Generate a Link Controller instance
func NewLinkController() LinkControllerInterface {
	return LinkController{
		LinkService:    service.NewLinkService(),
	}
}
