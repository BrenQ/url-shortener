package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"urlshortener/model"
	"urlshortener/utils"
)

// Generate a code with a short url
func (l LinkController) CreateLink(c * gin.Context) {
	var link model.UrlRequest

	if err := c.ShouldBindJSON(&link) ; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	// Validate an url
	if  ! utils.ValidUrl(link.Url)  {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Url invalid"})
		return
	}

	// Get a code to build a short url
	shortUrl, _ := utils.GenerateShortUrl()
	// Store an url
	_,err := l.LinkService.Create(model.Url{
		Original:  link.Url,
		Short:     shortUrl,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest , gin.H{
			"message" : "Bad request",
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"url" : utils.GetFullPath(c , shortUrl),
	})
}