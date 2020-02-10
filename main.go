package main

import (
	"github.com/joho/godotenv"
	r "urlshortener/router"
	s "urlshortener/server"
)

func main(){
	// Load .env variables
	_ = godotenv.Load()

	// dna := model.Url{Original: "google.com" , Short: "938x73i"}

	// Test database connection
	// it does not belong part of the original code
	//_ , err :=  db.NewConnection().Database("url").Collection("link").InsertOne(context.Background(),dna)
	// Set endpoints with a router
	newRouter := r.New()
	r := newRouter.Get()
	//Server run
	s.Run(r)
}
