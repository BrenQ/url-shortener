package utils

import (
	"net/url"
	"math/rand"
	"github.com/gin-gonic/gin"
	"time"
)


// Generate a random url. Temporal approach
func GenerateShortUrl() (string, error) {
	var letters = [] rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	const bytes = 6

	r := make([]rune , bytes)
	// No cache results
	rand.Seed(time.Now().UnixNano())
	for i := range r  {
		r[i] = letters[rand.Intn(len(letters))]
	}

	return string(r), nil
}
// Get a url full name
func GetFullPath(c * gin.Context , code string)  string {
	return c.Request.Host + "/" + code
}

// Validate an link
func ValidUrl ( link string ) bool {
	url, err := url.Parse(link)

	return err == nil &&
		   url.Host != "" &&
		   (url.Scheme =="https" ||
		   url.Scheme == "http")
}
