package web

type OrderProductResponse struct {
	Id        int `json:"id"`
	OrderId   int `json:"OrderId"`
	ProductId int `json:"ProductId"`
	Qty       int `json:"Qty"`
	Price     int `json:"Price"`
	Amount    int `json:"Amount"`
}
