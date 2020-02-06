package main

import (
	"context"
	"urlshortener/model"
	r "urlshortener/router"
	s "urlshortener/server"
	"github.com/joho/godotenv"
	"log"
	db "urlshortener/database"
)

func main(){

	_ = godotenv.Load()

	dna := model.Url{Original: "google.com" , Short: "938x73i"}

	// Test database connection
	_ , err :=  db.GetConnection().Database("url").Collection("link").InsertOne(context.Background(),dna)

	if err != nil {
		log.Println(500, err)
	}
	// Set endpoints with a router
	router := r.Get()
	//Server run
	s.Run(router)
}
