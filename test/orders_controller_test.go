package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/go_rest_api_test")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxLifetime(10 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	OrdersRepository := repository.NewOrdersRepository()
	OrdersService := service.NewOrdersService(OrdersRepository, db, validate)
	OrdersController := controller.NewOrdersController(OrdersService)
	router := app.NewRouter(OrdersController)

	return middleware.NewAuthMiddleware(router)

}

func truncateOrders(db *sql.DB) {
	db.Exec("TRUNCATE Orders")

}

func TestCreateOrdersSucces(t *testing.T) {
	db := setupTestDB()
	truncateOrders(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : "Gadget"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/apo/Orders", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Gadget", responseBody["date"].(map[string]interface{})["name"])
}

func TestCreateOrdersFailed(t *testing.T) {
	db := setupTestDB()
	truncateOrders(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/Orders", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])

}

func TestUpdateOrdersRequest(t *testing.T) {
	db := setupTestDB()
	truncateOrders(db)

	tx, _ := db.Begin()
	OrdersRepository := repository.NewOrdersRepository()
	Orders := OrdersRepository.Save(context.Background(), tx, domain.Orders{
		CustomerId: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : "Gadget"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/Orders/"+strconv.Itoa(Orders.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, Orders.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])
}

func TestUpdateOrdersFailed(t *testing.T) {
	db := setupTestDB()
	truncateOrders(db)

	tx, _ := db.Begin()
	OrdersRepository := repository.NewOrderRepository()
	Orders := OrdersRepository.Save(context.Background(), tx, domain.Orders{
		CustomerId: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : " "}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/Orders/"+strconv.Itoa(Orders.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])

}

func TestGetOrdersSucces(t *testing.T) {
	db := setupTestDB()
	truncateOrders(db)

	tx, _ := db.Begin()
	OrdersRepository := repository.NewOrdersRepository()
	Orders := OrdersRepository.Save(context.Background(), tx, domain.Orders{
		CustomerId: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/Orders/"+strconv.Itoa(Orders.Id), nil)
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, Orders.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, Orders.CustomerId, responseBody["data"].(map[string]interface{})["name"])
}

func TestGetOrdersFailed(t *testing.T) {
	db := setupTestDB()
	truncateOrders(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/Orders/404", nil)
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])

}

func TestDeleteOrdersSucces(t *testing.T) {
	db := setupTestDB()
	truncateOrders(db)

	tx, _ := db.Begin()
	OrdersRepository = repository.NewOrdersRepository()
	Orders := OrdersRepositorySave(context.Background(), tx, domain.Orders{
		CustomerId: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/Orders/"+strconv.Itoa(Orders.Id), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

}

func TestDeleteOrdersFailed(t *testing.T) {
	db := setupTestDB()
	truncateOrders(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/Orders/404", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])

}

func TestListOrdersSucces(t *testing.T) {
	db := setupTestDB()
	truncateOrders(db)

	tx, _ := db.Begin()
	OrdersRepository := repository.NewOrdersRepository()
	Orders1 := OrdersRepository.Save(context.Background(), tx, domain.Orders{
		CustomerId: "Gadget",
	})
	Orders2 := OrdersRepository.Save(context.Background(), tx, domain.Orders{
		CustomerId: "Computer",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/Orders", nil)
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

	fmt.Println(responseBody)

	fmt.Println(responseBody)

	var Orders = responseBody["data"].([]interface{})

	OrdersResponse1 := Orders[0].(map[string]interface{})
	OrderResponse2 := Orders[1].(map[string]interface{})

	assert.Equal(t, Orders1.Id, int(OrdersResponse1["id"].(float64)))
	assert.Equal(t, Orders1.CustomerId, OrdersResponse1["OrdersResponse"])

	assert.Equal(t, Orders2.Id, int(OrderResponse2["id"].(float64)))
	assert.Equal(t, Orders2.CustomerId, OrderResponse2["OrderResponse"])
}

func TestUnautorized(t *testing.T) {
	db := setupTestDB()
	truncateOrders(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/Orders", nil)
	request.Header.Add("X-API-Key", "SALAH")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 401, int(responseBody["code"].(float64)))
	assert.Equal(t, "UNAUTORIZED", responseBody["status"])
}
