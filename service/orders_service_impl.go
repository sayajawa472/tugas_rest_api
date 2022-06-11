package service

import (
	"context"
	"database/sql"
)

type OrdersServiceImpl struct {
	OrdersRepository repository.OrdersRepository
	DB               *sql.DB
	Validate         *validator.Validate
}

func NewOrdersService(OrdersRepository repository.OrdersRepository, DB *sql.DB, validate *validator.Validate) OrdersService {
	return &OrdersServiceImpl{
		OrdersRepository: OrdersRepository,
		DB:               DB,
		Validate:         validate,
	}

}

func (service *OrdersServiceImpl) Create(ctx context.Context, request web.OrdersCreateRequest) web.OrdersResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	Orders := domain.Orders{
		CustomerId:  request.CustomerId,
		TotalAmount: request.TotalAmount,
	}
	Orders = service.OrdersRepository.Save(ctx, tx, Orders)

	return helper.ToOrdersResponse(Orders)
}

func (service *OrdersServiceImpl) Update(ctx context.Context, request web.OrdersUpdateRequest) web.OrdersResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	Orders, err := service.OrdersRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	Orders.CustomerId = request.CustomerId
	Orders.TotalAmount = request.TotalAmount
	Orders = service.OrdersRepository.Update(ctx, tx, Orders)

	return helper.ToOrdersResponse(Orders)

}

func (service *OrdersServiceImpl) Delete(ctx context.Context, OrdersId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	Orders, err := service.OrdersRepository.FindById(ctx, tx, OrdersId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.OrdersRepository.Delete(ctx, tx, Orders)
}

func (service *OrdersServiceImpl) FindById(ctx context.Context, OrdersId int) web.OrdersResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	Orders, err := service.OrdersRepository.FindById(ctx, tx, OrdersId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToOrdersResponse(Orders)
}

func (service *OrdersServiceImpl) FindByAll(ctx context.Context) []web.OrdersResponse {

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	Orders := service.OrdersRepository.FindByAll(ctx, tx)

	return helper.ToOrdersResponses(Orders)

}
