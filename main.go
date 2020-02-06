package main

import (
	"context"
	r "urlshortener/router"
	s "urlshortener/server"
	"github.com/joho/godotenv"
	"log"
	db "urlshortener/database"
)

func main(){

	type Sample struct {
		Test string `json:"Dna,omitempty"`
	}

	_ = godotenv.Load()

	dna := Sample{Test: "test"}

	// Test database connection
	_ , err :=  db.GetConnection().Database("url").Collection("link").InsertOne(context.Background(),dna)

	if err != nil {
		log.Println(500, err)
	}
	// Set endpoints with a router
	router := r.Get()
	r.Load(router)

	if err != nil {
		log.Fatal(err)
	}

	server := s.New(router.Router)
	log.Fatal(server.ListenAndServe())

}
