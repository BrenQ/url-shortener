package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"urlshortener/model"
)
// Redirect to a link from a previously generated code
func (l LinkController) GetLink (c * gin.Context) {
	var link model.CodeRequest

	if err := c.ShouldBindUri(&link) ; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "We could not process url info" , "error": err.Error()})
		return
	}

	res, err := l.LinkService.Get(link.Code)

	if err != nil {
		c.JSON(http.StatusBadRequest , gin.H{"message": "Code not registered"})
	}

	 c.Redirect(301, res.Original)
}
