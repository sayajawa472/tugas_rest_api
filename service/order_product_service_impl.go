package service

import (
	"context"
	"database/sql"
)

type OrderProductServiceImpl struct {
	OrderProductRepository repository.OrderProductRepository
	DB                     *sql.DB
	Validate               *validator.Validate
}

func NewOrderProductService(OrderProductRepository repository.OrderProductRepository, DB *sql.DB, validate *validator.Validate) OrderProductService {
	return &OrderProductServiceImpl{
		OrderProductRepository: OrderProductRepository,
		DB:                     DB,
		Validate:               validate,
	}

}

func (service *OrderProductServiceImpl) Create(ctx context.Context, request web.OrderProductCreateRequest) web.OrderProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	OrderProduct := domain.OrderProduct{
		OrderId:   request.OrderId,
		ProductId: request.ProductId,
		Qty:       request.Qty,
		Price:     request.Price,
		Amount:    request.Amount,
	}
	OrderProduct = service.OrderProductRepository.Save(ctx, tx, OrderProduct)

	return helper.ToOrderProductResponse(OrderProduct)
}

func (service *OrderProductServiceImpl) Update(ctx context.Context, request web.OrderProductUpdateRequest) web.OrderProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	OrderProduct, err := service.OrderProductRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	OrderProduct.OrderId = request.OrderId
	OrderProduct.ProductId = request.ProductId
	OrderProduct.Qty = request.Qty
	OrderProduct.Price = request.Price
	OrderProduct.Amount = request.Amount
	OrderProduct = service.OrderProductRepository.Update(ctx, tx, OrderProduct)

	return helper.ToOrderProductResponse(OrderProduct)

}

func (service *OrderProductServiceImpl) Delete(ctx context.Context, OrderProductId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	OrderProduct, err := service.OrderProductRepository.FindById(ctx, tx, OrderProductId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.OrderProductRepository.Delete(ctx, tx, OrderProduct)
}

func (service *OrderProductServiceImpl) FindById(ctx context.Context, OrderProdukId int) web.OrderProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	OrderProduct, err := service.OrderProductRepository.FindById(ctx, tx, OrderProdukId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToOrderProductResponse(OrderProduct)
}

func (service *OrderProductServiceImpl) FindByAll(ctx context.Context) []web.OrderProductResponse {

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.OrderProductRepository.FindByAll(ctx, tx)

	return helper.ToOrderProductResponses(categories)

}
