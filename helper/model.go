package helper

import (
	"golang_rest_api/model/domain"
	"golang_rest_api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Nama: category.Nama,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}

func ToCustomerResponse(customer domain.Customer) web.CustomerResponse {
	return web.CustomerResponse{
		Id:          customer.Id,
		Nama:        customer.Name,
		Address:     customer.Address,
		Email:       customer.Email,
		PhoneNumber: customer.PhoneNumber,
	}
}

func ToCustomerResponses(Customers []domain.Customer) []web.CustomerResponse {
	var CustomerResponses []web.CustomerResponse
	for _, Customer := range Customers {
		CustomerResponses = append(CustomerResponses, ToCustomerResponse(Customer))
	}
	return CustomerResponses
}

func ToOrderProductResponse(op domain.OrderProduct) web.OrderProductResponse {
	return web.OrderProductResponse{
		Id:        op.Id,
		OrderId:   op.OrderId,
		ProductId: op.ProductId,
		Qty:       op.Qty,
		Price:     op.Price,
		Amount:    op.Amount,
	}
}

func ToOrderProductResponses(op []domain.OrderProduct) []web.OrderProductResponse {
	var responses []web.OrderProductResponse
	for _, OrderProduct := range op {
		responses = append(responses, ToOrderProductResponse(OrderProduct))
	}
	return responses
}

func ToOrdersResponse(Orders domain.Orders) web.OrdersResponse {
	return web.OrdersResponse{
		Id:          Orders.Id,
		CustomerId:  Orders.CustomerId,
		TotalAmount: Orders.TotalAmount,
	}
}

func ToOrdersResponses(Orders []domain.Orders) []web.OrdersResponse {
	var OrdersResponses []web.OrdersResponse
	for _, Orders := range Orders {
		OrdersResponses = append(OrdersResponses, ToOrdersResponse(Orders))
	}
	return OrdersResponses
}

func ToProductResponse(Product domain.Product) web.ProductResponse {
	return web.ProductResponse{
		Id:         Product.Id,
		Nama:       Product.Name,
		Price:      Product.CategoryId,
		CategoryId: Product.CategoryId,
	}
}
func ToProduct(p web.ProductResponse) domain.Product {
	return domain.Product{
		Id:         p.Id,
		Name:       p.Nama,
		Price:      p.Price,
		CategoryId: p.CategoryId,
	}

}

func ToProductResponses(Product []domain.Product) []web.ProductResponse {
	var ProductResponses []web.ProductResponse
	for _, product := range Product {
		ProductResponses = append(ProductResponses, ToProductResponse(product))
	}
	return ProductResponses
}
