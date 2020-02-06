package router

import (
	"github.com/gorilla/mux"
	"urlshortener/controller"
)


type Router struct {
	* mux.Router
}


func Get() *Router {
	router := mux.NewRouter()
	return &Router{router}
}

func Load( r * Router)  {

	v := r.Router.PathPrefix("/api/v1").Subrouter()

	v.HandleFunc("/links" , controller.CreateLink).Methods("POST")
	v.HandleFunc("/links/{id}" , controller.GetLink).Methods("GET")
}
