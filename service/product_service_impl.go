package service

import (
	"context"
	"database/sql"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewProductService(ProductRepository repository.ProductRepository, DB *sql.DB, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository: ProductRepository,
		DB:                DB,
		Validate:          validate,
	}

}

func (service *ProductServiceImpl) Create(ctx context.Context, request web.ProductCreateRequest) web.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	Product := domain.Product{
		Name:       request.Name,
		Price:      request.Price,
		CategoryId: request.CategoryId,
	}
	Product = service.ProductRepository.Save(ctx, tx, Product)

	return helper.ToProductResponse(Product)
}

func (service *ProductServiceImpl) Update(ctx context.Context, request web.ProductUpdateRequest) web.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	ProductsResponse, err := service.ProductRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	product := helper.ToProduct(ProductsResponse)
	product.Name = request.Nama
	product.Price = request.Price
	product.CategoryId = request.CategoryId
	product = service.ProductRepository.Update(ctx, tx, product)

	return helper.ToProductResponse(product)

}

func (service *ProductServiceImpl) Delete(ctx context.Context, ProductId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	productResponse, err := service.ProductRepository.FindById(ctx, tx, ProductId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	product := helper.ToProduct(productResponse)
	service.ProductRepository.Delete(ctx, tx, product)
}

func (service *ProductServiceImpl) FindById(ctx context.Context, ProductId int) web.ProductResponse {
	logrus.Info("product service start")
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	Product, err := service.ProductRepository.FindById(ctx, tx, ProductId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	logrus.Info("product service ended")
	return Product
}

func (service *ProductServiceImpl) FindByAll(ctx context.Context) []web.ProductResponse {
	logrus.Info("Product service find by all start")
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	Products := service.ProductRepository.FindByAll(ctx, tx)
	//return helper.ToProductResponses(Product)
	logrus.Info("product service ended")
	return Products

}
