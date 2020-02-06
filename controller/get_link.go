package controller

import (
	"github.com/gin-gonic/gin"
)

func GetLink (c * gin.Context) {
	c.JSON(200, gin.H{"message" : "Get link"})
}
