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

type OrderProductControllerImpl struct {
	OrderProductService service.OrderProductService
}

func NewOrderProductController(orderproductService service.OrderProductService) OrderProductController {
	return &OrderProductControllerImpl{
		orderproductService,
	}

}
func (controller *OrderProductControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	OrderProductCreateRequest := web.OrderProductCreateRequest{}
	helper.ReadFromRequestBody(request, &OrderProductCreateRequest)

	OrderProductResponse := controller.OrderProductService.Create(request.Context(), OrderProductCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   OrderProductResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrderProductControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	OrderProductUpdateRequest := web.OrderProductUpdateRequest{}
	helper.ReadFromRequestBody(request, &OrderProductUpdateRequest)

	OrderProductId := params.ByName("OrderProductId")
	id, err := strconv.Atoi(OrderProductId)
	helper.PanicIfError(err)

	OrderProductUpdateRequest.Id = id

	OrderProductResponse := controller.OrderProductService.Update(request.Context(), OrderProductUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   OrderProductResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrderProductControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	OrderProductId := params.ByName("OrderProductId")
	id, err := strconv.Atoi(OrderProductId)
	helper.PanicIfError(err)

	controller.OrderProductService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrderProductControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	logrus.Info("order product controller find by id start")
	OrderProductId := params.ByName("OrderProductId")
	id, err := strconv.Atoi(OrderProductId)
	helper.PanicIfError(err)

	OrderProductResponse := controller.OrderProductService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   OrderProductResponse,
	}

	logrus.Info("order product controller find by id ended")
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller OrderProductControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	logrus.Info("order product find all start")
	OrderProductResponses := controller.OrderProductService.FindByAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   OrderProductResponses,
	}

	logrus.Info("order product find all ended")
	helper.WriteToResponseBody(writer, webResponse)
}
