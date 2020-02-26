package main

import (
	c "urlshortener/config"
	d "urlshortener/database"
	r "urlshortener/router"
	s "urlshortener/server"
)

func main(){
	// Load .env variables
	cfg := c.NewConfig()
	cfg.Start()

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
