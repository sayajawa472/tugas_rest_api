package web

type CategoryCreateRequest struct {
	Nama string `validate:"required,min=1,max=100" json:"nama"`
}
