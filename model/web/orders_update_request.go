package web

type OrdersUpdateRequest struct {
	Id          int `validate:"required"`
	CustomerId  int `validate:"required,max=200,min=1" json:"Customer_id"`
	TotalAmount int `validate:"required,max=200,min=1" json:"Total_amount"`
}
