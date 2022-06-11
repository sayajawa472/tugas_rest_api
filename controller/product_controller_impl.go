package controller

import (
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"golang_rest_api/helper"
	"golang_rest_api/model/web"
	"golang_rest_api/service"
	"net/http"
	"strconv"
)

type ProductControllerImpl struct {
	ProductService service.ProductService
}

func NewProductController(ProductService service.ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: ProductService,
	}
}

func (controller *ProductControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ProductCreateRequest := web.ProductCreateRequest{}
	helper.ReadFromRequestBody(request, &ProductCreateRequest)

	ProductResponse := controller.ProductService.Create(request.Context(), ProductCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   ProductResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ProductUpdateRequest := web.ProductUpdateRequest{}
	helper.ReadFromRequestBody(request, &ProductUpdateRequest)

	ProductId := params.ByName("ProductId")
	id, err := strconv.Atoi(ProductId)
	helper.PanicIfError(err)

	ProductUpdateRequest.Id = id

	ProductResponse := controller.ProductService.Update(request.Context(), ProductUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   ProductResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ProductId := params.ByName("ProductId")
	id, err := strconv.Atoi(ProductId)
	helper.PanicIfError(err)

	controller.ProductService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	logrus.Info("product controller find by id start")
	ProductId := params.ByName("ProductId")
	id, err := strconv.Atoi(ProductId)
	helper.PanicIfError(err)

	ProductResponse := controller.ProductService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   ProductResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
	logrus.Info("find by id ended")
}

func (controller *ProductControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	logrus.Info("product controller find by all start")
	ProductResponses := controller.ProductService.FindByAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   ProductResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
	logrus.Info("product controller ended")
}
