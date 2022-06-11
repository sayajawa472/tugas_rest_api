package service

import "context"

type OrdersService interface {
	Create(ctx context.Context, request web.OrdersCreateRequest) web.OrdersResponse
	Update(ctx context.Context, request web.OrdersUpdateRequest) web.OrdersResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) web.OrdersResponse
	FindByAll(ctx context.Context) []web.OrdersResponse
}
