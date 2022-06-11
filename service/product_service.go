package service

import "context"

type ProductService interface {
	Create(ctx context.Context, request web.ProductCreateRequest) web.ProductResponse
	Update(ctx context.Context, request web.ProductUpdateRequest) web.ProductResponse
	Delete(ctx context.Context, ProductId int)
	FindById(ctx context.Context, ProductId int) web.ProductResponse
	FindByAll(ctx context.Context) []web.ProductResponse
}
