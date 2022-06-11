package service

import (
	"context"
	"database/sql"
)

type CustomerServiceImpl struct {
	CustomerRepository repository.CustomerRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCustomerService(CustomerRepository repository.CustomerRepository, DB *sql.DB, validate *validator.Validate) CustomerService {
	return &CustomerServiceImpl{
		CustomerRepository: CustomerRepository,
		DB:                 DB,
		Validate:           validate,
	}

}

func (service *CustomerServiceImpl) Create(ctx context.Context, request web.CustomerCreateRequest) web.CustomerResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	Customer := domain.Customer{
		Name:        request.Nama,
		Address:     request.Address,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
	}
	Customer = service.CustomerRepository.Save(ctx, tx, Customer)

	return helper.ToCustomerResponse(Customer)
}

func (service *CustomerServiceImpl) Update(ctx context.Context, request web.CustomerUpdateRequest) web.CustomerResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customer, err := service.CustomerRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	customer.Name = request.Nama
	customer.Address = request.Address
	customer.Email = request.Email
	customer.PhoneNumber = request.PhoneNumber
	customer = service.CustomerRepository.Update(ctx, tx, customer)

	return helper.ToCustomerResponse(customer)

}

func (service *CustomerServiceImpl) Delete(ctx context.Context, customerId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customer, err := service.CustomerRepository.FindById(ctx, tx, customerId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.CustomerRepository.Delete(ctx, tx, customer)
}

func (service *CustomerServiceImpl) FindById(ctx context.Context, customerId int) web.CustomerResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CustomerRepository.FindById(ctx, tx, customerId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToCustomerResponse(category)
}

func (service *CustomerServiceImpl) FindByAll(ctx context.Context) []web.CustomerResponse {

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CustomerRepository.FindByAll(ctx, tx)

	return helper.ToCustomerResponses(categories)

}
