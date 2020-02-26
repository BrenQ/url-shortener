package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	c "urlshortener/config"
)

// Run the server
func Run (router *gin.Engine) {
	log.Fatal(router.Run(GetAddr()))
}

// Address server
func GetAddr () string {
	var cfg  c.Config
	addr := fmt.Sprintf("%s:%s" , cfg.Get("SERVER_ADDR"), cfg.Get("SERVER_PORT"))
	return addr
}