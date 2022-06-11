package web

type ProductUpdateRequest struct {
	Id         int    `validate:"required"`
	Nama       string `validate:"required,max=200,min=1" json:"nama"`
	Price      int    `validate:"required" json:"Price"`
	CategoryId int    `validate:"required" json:"CategoryId"`
}
