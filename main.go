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
	_, err := d.Init()

	if err != nil {
		panic(err)
	}
	// Routes
	newRouter := r.New()
	r := newRouter.Get()
	//Server run
	s.Run(r)
}
