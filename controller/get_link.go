package controller

import (
	"github.com/gin-gonic/gin"
)

func (l LinkController) GetLink (c * gin.Context) {
	c.JSON(200, gin.H{"message" : "Get link"})
}
