package controllers

import (
	"encoding/json"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("ola")
}
