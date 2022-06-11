package web

type OrderProductCreateRequest struct {
	OrderId   int `validate:"required,min=1,max=100" json:"OrderId"`
	ProductId int `validate:"required,min=1,max=100" json:"ProductId"`
	Qty       int `validate:"required,min=1,max=100" json:"Qty"`
	Price     int `validate:"required,min=1,max=100" json:"Price"`
	Amount    int `validate:"required,min=1,max=100" json:"Amount"`
}
