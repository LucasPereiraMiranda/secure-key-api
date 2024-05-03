package controllers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"
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

func GeneratePassword(w http.ResponseWriter, r *http.Request) {
	numDigitsStr := r.URL.Query().Get("num_digits")
	hasSpecialKeysStr := r.URL.Query().Get("has_special_keys")
	hasNumbersKeyStr := r.URL.Query().Get("has_numbers_key")
	upperCaseKeyStr := r.URL.Query().Get("upper_case_key")

	numDigits, err := strconv.Atoi(numDigitsStr)
	if err != nil {
		respondWithError(w, "Invalid number of digits")
		return
	}

	hasSpecialKeys, err := strconv.ParseBool(hasSpecialKeysStr)
	if err != nil {
		respondWithError(w, "Invalid has_special_keys parameter")
		return
	}

	hasNumbersKey, err := strconv.ParseBool(hasNumbersKeyStr)
	if err != nil {
		respondWithError(w, "Invalid has_numbers_key parameter")
		return
	}

	upperCaseFlag, err := strconv.ParseBool(upperCaseKeyStr)
	if err != nil {
		respondWithError(w, "Invalid upper_case_flag parameter")
		return
	}

	var chars []rune
	if hasSpecialKeys {
		chars = append(chars, []rune("!@#$%^&*()_+")[0:10]...)
	}
	if hasNumbersKey {
		chars = append(chars, []rune("0123456789")[0:10]...)
	}
	chars = append(chars, []rune("abcdefghijklmnopqrstuvwxyz")...)
	if upperCaseFlag {
		chars = append(chars, []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")...)
	}

	randGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))

	password := make([]rune, numDigits)
	for i := range password {
		password[i] = chars[randGenerator.Intn(len(chars))]
	}

	response := PasswordResponse{
		Password: string(password),
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		respondWithError(w, "Error encoding JSON")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func respondWithError(w http.ResponseWriter, errMsg string) {
	errResponse := ErrorResponse{Error: errMsg}
	jsonResponse, err := json.Marshal(errResponse)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(jsonResponse)
}
