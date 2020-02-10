package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

// Run the server
func Run (router *gin.Engine) {
	log.Fatal(router.Run(GetAddr()))
}

// Address server
func GetAddr () string {
	addr := fmt.Sprintf("%s:%s" , os.Getenv("SERVER_ADDR"), os.Getenv("SERVER_PORT"))
	return addr
}