package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)


func Run (router *gin.Engine) {
	log.Fatal(router.Run(GetAddr()))
}

func GetAddr () string{
	addr := fmt.Sprintf("%s:%s" , os.Getenv("SERVER_ADDR"), os.Getenv("SERVER_PORT"))
	return addr
}