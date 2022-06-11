package web

type OrdersCreateRequest struct {
	CustomerId  int `validate:"required,min=1,max=100" json:"CustomerId"`
	TotalAmount int `validate:"required,min=1,max=100" json:"TotalAmount"`
}
