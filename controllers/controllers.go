package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"status": "healthy"}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error on encode response", http.StatusInternalServerError)
		return
	}
	w.Write(responseJSON)
}
func GeneratePassword(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Start Page")
}
