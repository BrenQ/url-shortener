package controller

import (
	"encoding/json"
	"net/http"
)

func CreateLink( w http.ResponseWriter , r * http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode("Create link")
}