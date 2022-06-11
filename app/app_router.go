package app

import (
	"github.com/julienschmidt/httprouter"
	"golang_rest_api/controller"
	"golang_rest_api/exception"
)

func NewRouter(categoryController controller.CategoryController, customerController controller.CustomerController, ProductController controller.ProductController, OrderProductController controller.OrderProductController, OrdersController controller.OrdersController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.GET("/api/customers", customerController.FindAll)
	router.GET("/api/customers/:customersId", customerController.FindById)
	router.POST("/api/customers", customerController.Create)
	router.PUT("/api/customers/:customerId", customerController.Update)
	router.DELETE("/api/customers/:customerId", customerController.Delete)

	router.GET("/api/OrderProduct", OrderProductController.FindAll)
	router.GET("/api/OrderProduct/:OrderProductId", OrderProductController.FindById)
	router.POST("/api/OrderProduct", OrderProductController.Create)
	router.PUT("/api/OrderProduct/:OrderProductId", OrderProductController.Update)
	router.DELETE("/api/OrderProduct/:OrderProductId", OrderProductController.Delete)

	router.GET("/api/Orders", OrdersController.FindAll)
	router.GET("/api/Orders/:OrdersId", OrdersController.FindById)
	router.POST("/api/Orders", OrdersController.Create)
	router.PUT("/api/Orders/:OrdersId", OrdersController.Update)
	router.DELETE("/api/Orders/:OrdersId", OrdersController.Delete)

	router.GET("/api/Product", ProductController.FindAll)
	router.GET("/api/Product/:ProductId", ProductController.FindById)
	router.POST("/api/Product", ProductController.Create)
	router.PUT("/api/Product/:ProductId", ProductController.Update)
	router.DELETE("/api/Product/:ProductId", ProductController.Delete)

	router.PanicHandler = exception.ErrorHandler
	return router

}
