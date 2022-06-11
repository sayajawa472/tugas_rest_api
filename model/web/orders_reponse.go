package web

type OrdersResponse struct {
	Id          int `json:"id"`
	CustomerId  int `json:"CustomerId"`
	TotalAmount int `json:"TotalAmount"`
}
