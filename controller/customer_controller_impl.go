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

type CustomerControllerImpl struct {
	CustomerService service.CustomerService
}

func NewCustomerController(customerService service.CustomerService) CustomerController {
	return &CustomerControllerImpl{
		CustomerService: customerService,
	}
}

func (controller *CustomerControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	customerCreateRequest := web.CustomerCreateRequest{}
	helper.ReadFromRequestBody(request, &customerCreateRequest)

	customerResponse := controller.CustomerService.Create(request.Context(), customerCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CustomerControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	customerUpdateRequest := web.CustomerUpdateRequest{}
	helper.ReadFromRequestBody(request, &customerUpdateRequest)

	customerId := params.ByName("customerId")
	id, err := strconv.Atoi(customerId)
	helper.PanicIfError(err)

	customerUpdateRequest.Id = id

	customerResponse := controller.CustomerService.Update(request.Context(), customerUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CustomerControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	customerId := params.ByName("customerId")
	id, err := strconv.Atoi(customerId)
	helper.PanicIfError(err)

	controller.CustomerService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CustomerControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	logrus.Info("customer controller find by id start")
	customerId := params.ByName("customerId")
	id, err := strconv.Atoi(customerId)
	helper.PanicIfError(err)

	customerResponse := controller.CustomerService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerResponse,
	}

	logrus.Info("customer controller find by id ended")
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CustomerControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	logrus.Info("customer controller find all start")
	customerResponses := controller.CustomerService.FindByAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerResponses,
	}

	logrus.Info("customer controller find all ended")
	helper.WriteToResponseBody(writer, webResponse)
}
