package models

// Address informações de entrada do endereço
type Address struct {
	Street     string `json:"street,omitempty"`
	Number     string `json:"number,omitempty"`
	Complement string `json:"complement,omitempty"`
	ZipCode    string `json:"zipcode,omitempty"`
	City       string `json:"city,omitempty"`
	District   string `json:"district,omitempty"`
	StateCode  string `json:"state_code,omitempty"`
}
