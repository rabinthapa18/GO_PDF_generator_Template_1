package models

type RequestBody struct {
	Template    string      `json:"template"`
	Definitions Definitions `json:"definitions"`
	Values      Values      `json:"values"`
}
