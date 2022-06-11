package service

import "context"

type CustomerService interface {
	Create(ctx context.Context, request web.CustomerCreateRequest) web.CustomerResponse
	Update(ctx context.Context, request web.CustomerUpdateRequest) web.CustomerResponse
	Delete(ctx context.Context, customerId int)
	FindById(ctx context.Context, customerId int) web.CustomerResponse
	FindByAll(ctx context.Context) []web.CustomerResponse
}
