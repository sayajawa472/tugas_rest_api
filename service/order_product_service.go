package service

import "context"

type OrderProductService interface {
	Create(ctx context.Context, OrderProduct web.OrderProductCreateRequest) web.OrderProductResponse
	Update(ctx context.Context, OrderProduct web.OrderProductUpdateRequest) web.OrderProductResponse
	Delete(ctx context.Context, OrderProductId int)
	FindById(ctx context.Context, OrderProductId int) web.OrderProductResponse
	FindByAll(ctx context.Context) []web.OrderProductResponse
}
