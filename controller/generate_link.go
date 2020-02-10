package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"urlshortener/model"
	"urlshortener/utils"
)

func (l LinkController) CreateLink(c * gin.Context) {
	var link *model.UrlRequest
	
	if err := c.ShouldBindJSON(&link) ; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "We could not process url info" , "error": err.Error()})
		return
	}
	
	shortUrl := utils.GenerateShortUrl()

	l.LinkService.Create(model.Url{
		Original: link.Url,
		Short:     shortUrl,
		CreatedAt: time.Time{},
	})

	c.JSON(200, link)
}