package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestGenerateShortUrl_CheckReturnAnyValue(t *testing.T) {
	url ,_ := GenerateShortUrl()

	if url == "" {
		t.Errorf("Url is empty")
	}
}

func TestGenerateShortUrl_CheckReturnRandomValue(t *testing.T) {
	ass := assert.New(t)
	url1 ,_ := GenerateShortUrl()
	url2 ,_ := GenerateShortUrl()
	
	ass.NotEqual(url1,url2)
}

func TestValidUrl_CheckInvalidUrl(t *testing.T) {
	url := "http:/google"

	if ValidUrl(url) {
		t.Errorf("Url is considered valid")
	}
}

func TestValidUrl_CheckInvalidUrlProtocol(t *testing.T) {
	url := "httpa://www.google.com"

	if ValidUrl(url) {
		t.Errorf("Url is considered valid")
	}
}

func TestValidUrl_CheckValidUrl(t *testing.T) {
	url := "http://google.com"

	if ! ValidUrl(url) {
		t.Errorf("Url is considered valid")
	}
}

