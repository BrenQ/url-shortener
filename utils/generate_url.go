package utils

import "github.com/gin-gonic/gin"

func GenerateShortUrl() string {
	return "9wTAM5"
}

func GetFullPath(c * gin.Context , code string)  string {
	return c.Request.Host + "/" + code
}
