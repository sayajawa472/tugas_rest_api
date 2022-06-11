package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"golang_rest_api/helper"
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

func truncateCustomer(db *sql.DB) {
	db.Exec("TRUNCATE customer")
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	CustomerRepository := repository.NewCustomerRepository()
	CustomerService := service.NewCustomerService(CustomerRepository, db, validate)
	CustomerController := controller.NewCustomerController(CustomerService)
	router := app.NewRouter(CustomerController)

	return middleware.NewAuthMiddleware(router)

}

func TestCreateCustomerSucces(t *testing.T) {
	db := setupTestDB()
	truncateCustomer(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : "Gadget"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/apo/customer", requestBody)
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

func TestCustomerFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/customer", requestBody)
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

func TestUpdateCustomerRequest(t *testing.T) {
	db := setupTestDB()
	truncateCustomer(db)

	tx, _ := db.Begin()
	CustomerRepository := repository.NewCustomerRepository()
	Customer := CustomerRepository.Save(context.Background(), tx, domain.Customer{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : "Gadget"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/customer/"+strconv.Itoa(Customer.Id), requestBody)
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
	assert.Equal(t, Customer.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])
}

func TestUpdateCustomerCreate(t *testing.T) {
	db := setupTestDB()
	truncateCustomer(db)

	tx, _ := db.Begin()
	CustomerRepository := repository.NewCustomerRepository()
	Customer := CustomerRepository.Save(context.Background(), tx, domain.Customer{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : " "}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/customer/"+strconv.Itoa(Customer.Id), requestBody)
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

func TestCustomerSucces(t *testing.T) {
	db := setupTestDB()
	truncateCustomer(db)

	tx, _ := db.Begin()
	CustomerRepository := repository.NewCustomerRepository()
	Customer := CustomerRepository.Save(context.Background(), tx, domain.Customer{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/customer/"+strconv.Itoa(Customer.Id), nil)
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
	assert.Equal(t, Customer.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, Customer.Nama, responseBody["data"].(map[string]interface{})["nama"])
}

func TestGetCustomerFailed(t *testing.T) {
	db := setupTestDB()
	truncateCustomer(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/customer/404", nil)
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

func TestDeleteCustomerSucces(t *testing.T) {
	db := setupTestDB()
	truncateCustomer(db)

	tx, _ := db.Begin()
	CustomerRepository := repository.NewCustomerRepository()
	Customer := CustomerRepository.Save(context.Background(), tx, domain.Customer{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/customer/"+strconv.Itoa(Customer.Id), nil)
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

func TestDeleteCustomerFailed(t *testing.T) {
	db := setupTestDB()
	truncateCustomer(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/customer/404", nil)
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

func TestListCustomerSucces(t *testing.T) {
	db := setupTestDB()
	truncateCustomer(db)

	tx, _ := db.Begin()
	CustomerRepository := repository.NewCustomerRepository()
	Customer1 := CustomerRepository.Save(context.Background(), tx, domain.Customer{
		Name: "Gadget",
	})
	Customer2 := CustomerRepository.Save(context.Background(), tx, domain.Customer{
		Name: "Computer",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/customer", nil)
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

	var categories = responseBody["data"].([]interface{})

	categoryResponse1 := categories[0].(map[string]interface{})
	categoryResponse2 := categories[1].(map[string]interface{})

	assert.Equal(t, Customer1.Id, int(categoryResponse1["id"].(float64)))
	assert.Equal(t, Customer2.Nama, categoryResponse1["nama"])

	assert.Equal(t, Customer2.Id, int(categoryResponse2["id"].(float64)))
	assert.Equal(t, Customer2.Nama, categoryResponse2["name"])
}

func TestUnautorized(t *testing.T) {
	db := setupTestDB()
	truncateCustomer(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/customer", nil)
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
