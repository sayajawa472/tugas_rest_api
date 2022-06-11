package web

type CustomerResponse struct {
	Id          int    `json:"id"`
	Nama        string `json:"nama"`
	Address     string `json:"Address"`
	Email       string `json:"Email"`
	PhoneNumber string `json:"PhoneNumber"`
}
