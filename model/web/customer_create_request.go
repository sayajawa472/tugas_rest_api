package web

type CustomerCreateRequest struct {
	Nama        string `validate:"required,min=1,max=100" json:"nama"`
	Address     string `validate:"required,min=1,max=100" json:"Address"`
	Email       string `validate:"required,min=1,max=100" json:"Email"`
	PhoneNumber string `validate:"required,min=1,max=100" json:"PhoneNumber"`
}
