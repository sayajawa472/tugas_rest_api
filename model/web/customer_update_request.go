package web

type CustomerUpdateRequest struct {
	Id          int    `validate:"required"`
	Nama        string `validate:"required,max=200,min=1" json:"nama"`
	Address     string `validate:"required,max=200,min=1" json:"Address"`
	Email       string `validate:"required,max=200,min=1" json:"Email"`
	PhoneNumber string `validate:"required,max=200,min=1" json:"PhoneNumber"`
}
