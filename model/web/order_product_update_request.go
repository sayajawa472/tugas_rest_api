package web

type OrderProductUpdateRequest struct {
	Id        int `validate:"required"`
	OrderId   int `validate:"required,max=200,min=1" json:"orderId"`
	ProductId int `validate:"required,max=200,min=1" json:"ProductId"`
	Qty       int `validate:"required,max=200,min=1" json:"Qty"`
	Price     int `validate:"required,max=200,min=1" json:"Price"`
	Amount    int `validate:"required,max=200,min=1" json:"Amount"`
}
