package test

/**
import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"go-rest-api/app"
	"go-rest-api/controller"
	"go-rest-api/helper"
	"go-rest-api/middleware"
	"go-rest-api/model/domain"
	"go-rest-api/repository"
	"go-rest-api/service"
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
	ProductRepository := repository.NewProductRepository()
	ProductService := service.NewProductService(ProductRepository, db, validate)
	ProductController := controller.NewProductController(Service)
	router := app.NewRouter(ProductController)

	return middleware.NewAuthMiddleware(router)

}

func truncateProduct(db *sql.DB) {
	db.Exec("TRUNCATE Product")

}

func TestCreateProductSucces(t *testing.T) {
	db := setupTestDB()
	truncateProduct(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : "Gadget"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/apo/Product", requestBody)
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

func TestCreateProductFailed(t *testing.T) {
	db := setupTestDB()
	truncateProduct(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/Product", requestBody)
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

func TestUpdateProductRequest(t *testing.T) {
	db := setupTestDB()
	truncateProduct(db)

	tx, _ := db.Begin()
	ProductRepository := repository.NewProductRepository()
	Product := ProductRepository.Save(context.Background(), tx, domain.Product{
		Nama: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : "Gadget"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/Product/"+strconv.Itoa(Product.Id), requestBody)
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
	assert.Equal(t, Product.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])
}

func TestUpdateProductFailed(t *testing.T) {
	db := setupTestDB()
	truncateProduct(db)

	tx, _ := db.Begin()
	ProductRepository := repository.NewProductRepository()
	Product := ProductRepository.Save(context.Background(), tx, domain.Product{
		Nama: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : " "}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/Product/"+strconv.Itoa(Product.Id), requestBody)
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

func TestGetProductSucces(t *testing.T) {
	db := setupTestDB()
	truncateProduct(db)

	tx, _ := db.Begin()
	ProductRepository := repository.NewCategoryRepository()
	Product := ProductRepository.Save(context.Background(), tx, domain.Product{
		Product: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/Product/"+strconv.Itoa(Product.Id), nil)
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
	assert.Equal(t, Product.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, Product.Nama, responseBody["data"].(map[string]interface{})["name"])
}

func TestGetProductFailed(t *testing.T) {
	db := setupTestDB()
	truncateProduct(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/Product/404", nil)
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

func TestDeleteProductSucces(t *testing.T) {
	db := setupTestDB()
	truncateProduct(db)

	tx, _ := db.Begin()
	ProductRepository := repository.NewProductRepository()
	Product := ProductRepository.Save(context.Background(), tx, domain.Product{
		Product: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/Product/"+strconv.Itoa(Product.Id), nil)
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

func TestDeleteProductFailed(t *testing.T) {
	db := setupTestDB()
	truncateProduct(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/Product/404", nil)
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

func TestListProductSucces(t *testing.T) {
	db := setupTestDB()
	truncateProduct(db)

	tx, _ := db.Begin()
	ProductRepository := repository.NewProductRepository()
	Product1 := ProductRepository.Save(context.Background(), tx, domain.Product{
		Nama: "Gadget",
	})
	Product2 := categoryRepository.Save(context.Background(), tx, domain.Product{
		Nama: "Computer",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/Product", nil)
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

	var categories = responseBody["data"].([]interface{})

	ProductResponse1 := categories[0].(map[string]interface{})
	categoryResponse2 := categories[1].(map[string]interface{})

	assert.Equal(t, Product1.Id, int(ProductResponse1["id"].(float64)))
	assert.Equal(t, Product2.Nama, ProductResponse1["name"])

	assert.Equal(t, Product2.Id, int(ProductResponse2["id"].(float64)))
	assert.Equal(t, Product2.Nama, ProductResponse2["name"])
}

func TestUnautorized(t *testing.T) {
	db := setupTestDB()
	truncateProduct(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/Product", nil)
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

*/
