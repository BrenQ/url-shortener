package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)


func New (router * mux.Router) *http.Server {
	return &http.Server{
		Handler: router,
		Addr: fmt.Sprintf("%s:%s" , os.Getenv("SERVER_ADDR"), os.Getenv("SERVER_PORT") ),
	}
}