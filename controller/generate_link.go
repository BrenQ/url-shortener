package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"urlshortener/model"
	"urlshortener/utils"
)

// Generate a code with a short url
func (l LinkController) CreateLink(c * gin.Context) {
	var link *model.UrlRequest

	if err := c.ShouldBindJSON(&link) ; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "We could not process url info" , "error": err.Error()})
		return
	}
	// Get a code to build a short url
	shortUrl := utils.GenerateShortUrl()
	// Store an url
	err := l.LinkService.Create(model.Url{
		Original:  link.Url,
		Short:     shortUrl,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest , gin.H{
			"message" : "Bad request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"url" : utils.GetFullPath(c , shortUrl),
	})
}