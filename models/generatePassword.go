package models

type GeneratePasswordRequest struct {
	NumDigits      int  `json:"num_digits"`
	HasSpecialKeys bool `json:"has_special_keys"`
	HasNumbersKey  bool `json:"has_numbers_key"`
	UpperCaseKey   bool `json:"upper_case_key"`
}

type PasswordResponse struct {
	Password string `json:"password"`
}
