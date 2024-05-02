package controllers

import (
	"encoding/json"
	"net/http"
)

type GeneratePasswordRequest struct {
	NumDigits      int  `json:"num_digits"`
	HasSpecialKeys bool `json:"has_special_keys"`
	HasNumbersKey  bool `json:"has_numbers_key"`
	UpperCaseKey   bool `json:"upper_case_key"`
}

type PasswordResponse struct {
	Password string `json:"password"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func Health(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"status": "healthy"}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error on encode response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}
