package utils

import (
	"math/rand"
	"github.com/gin-gonic/gin"
)


// Generate a random url. Temporal approach
func GenerateShortUrl() (string, error) {
	var letters = [] rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	const bytes = 6

	r := make([]rune , bytes)

	for i := range r  {
		r[i] = letters[rand.Intn(len(letters))]
	}

	return string(r), nil
}

func GetFullPath(c * gin.Context , code string)  string {
	return c.Request.Host + "/" + code
}
