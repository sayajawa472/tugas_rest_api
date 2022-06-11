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

type OrdersControllerImpl struct {
	OrdersService service.OrdersService
}

func NewOrdersController(OrdersService service.OrdersService) OrdersController {
	return &OrdersControllerImpl{
		OrdersService: OrdersService,
	}
}

func (orders *OrdersControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	OrdersCreateRequest := web.OrdersCreateRequest{}
	helper.ReadFromRequestBody(request, &OrdersCreateRequest)

	OrdersResponse := orders.OrdersService.Create(request.Context(), OrdersCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   OrdersResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (orders *OrdersControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	OrdersUpdateRequest := web.OrdersUpdateRequest{}
	helper.ReadFromRequestBody(request, &OrdersUpdateRequest)

	OrdersId := params.ByName("OrdersId")
	id, err := strconv.Atoi(OrdersId)
	helper.PanicIfError(err)

	OrdersUpdateRequest.Id = id

	OrdersResponse := orders.OrdersService.Update(request.Context(), OrdersUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   OrdersResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrdersControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	OrdersId := params.ByName("OrdersId")
	id, err := strconv.Atoi(OrdersId)
	helper.PanicIfError(err)

	controller.OrdersService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrdersControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	logrus.Info("orders controller find by id start")
	OrdersId := params.ByName("OrdersId")
	id, err := strconv.Atoi(OrdersId)
	helper.PanicIfError(err)

	OrdersResponse := controller.OrdersService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   OrdersResponse,
	}

	logrus.Info("orders controller find by id ended")
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrdersControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	logrus.Info("orders controller find all start")
	OrdersResponses := controller.OrdersService.FindByAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   OrdersResponses,
	}

	logrus.Info("orders controller find all ended")
	helper.WriteToResponseBody(writer, webResponse)
}
