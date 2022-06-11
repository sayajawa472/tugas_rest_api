package web

type ProductCreateRequest struct {
	Name       string `validate:"required,min=1,max=100" json:"nama"`
	Price      int    `validate:"required" json:"Price"`
	CategoryId int    `validate:"required" json:"CategoryId"`
}
