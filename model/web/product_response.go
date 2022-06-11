package web

type ProductResponse struct {
	Id           int    `json:"id"`
	Nama         string `json:"nama"`
	Price        int    `json:"price"`
	CategoryId   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
}
