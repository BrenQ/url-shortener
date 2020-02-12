package main

import (
	"github.com/joho/godotenv"
	d "urlshortener/database"
	r "urlshortener/router"
	s "urlshortener/server"
)

func main(){
	// Load .env variables
	_ = godotenv.Load()
	// Init database config
	d.Init()

	newRouter := r.New()
	// Get an engine instance with routes
	r := newRouter.Get()
	//Server run
	s.Run(r)
}
