package utils

import (
	"math/rand"
	"net/url"
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

// Validate an link
func ValidUrl ( link string ) bool {
	url, err := url.Parse(link)

	return err == nil &&
		   url.Host != "" &&
		   (url.Scheme =="https" ||
		   url.Scheme == "http")
}
