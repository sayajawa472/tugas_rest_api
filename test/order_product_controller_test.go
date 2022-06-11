package test

/**
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
	OrderProductRepository := repository.NewOrderProductRepository()
	OrderProductRepository := service.NewOrderProductService(OrderProductRepository, db, validate)
	OrderProductController := controller.NewOrderProductController(OrderProductService)
	router := app.NewRouter(OrderProductController)

	return middleware.NewAuthMiddleware(router)

}

func truncateOrderProduct(db *sql.DB) {
	db.Exec("TRUNCATE OrderProduct")

}

func TestCreateOrderProductSucces(t *testing.T) {
	db := setupTestDB()
	truncateOrderProduct(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : "Gadget"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/OrderProduct", requestBody)
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

func TestCreateOrderProductFailed(t *testing.T) {
	db := setupTestDB()
	truncateOrderProduct(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/OrderProduct", requestBody)
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

func TestUpdateOrderProductRequest(t *testing.T) {
	db := setupTestDB()
	truncateOrderProduct(db)

	tx, _ := db.Begin()
	OrderProductRepository := repository.NewOrderProductRepository()
	OrderProduct := OrderProductRepository.Save(context.Background(), tx, domain.OrderProduct{
		Nama: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : "Gadget"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/OrderProduct/"+strconv.Itoa(OrderProduct.Id), requestBody)
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
	assert.Equal(t, OrderProduct.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])
}

func TestUpdateOrderProductFailed(t *testing.T) {
	db := setupTestDB()
	truncateOrderProduct(db)

	tx, _ := db.Begin()
	OrderProductRepository := repository.NewOrderProductRepository()
	OrderProduct := OrderProductRepository.Save(context.Background(), tx, domain.OrderProduct{
		Nama: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : " "}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/OrderProduct/"+strconv.Itoa(OrderProduct.Id), requestBody)
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

func TestGetOrderProductSucces(t *testing.T) {
	db := setupTestDB()
	truncateOrderProduct(db)

	tx, _ := db.Begin()
	OrderProductRepository := repository.NewCategoryRepository()
	OrderProduct := OrderProductRepository.Save(context.Background(), tx, domain.OrderProduct{}
Nama: "Gadget",
})
tx.Commit()

router := setupRouter(db)

request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/OrderProduct/"+strconv.Itoa(OrderProduct.Id), nil)
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
assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
assert.Equal(t, category.Nama, responseBody["data"].(map[string]interface{})["name"])
}

func TestGetOrderProductFailed(t *testing.T) {
	db := setupTestDB()
	truncateOrderProduct(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/OrderProduct/404", nil)
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

func TestDeleteOrderProductSucces(t *testing.T) {
	db := setupTestDB()
	truncateOrderProduct(db)

	tx, _ := db.Begin()
	OrderProductRepository := repository.NewOrderProductRepository()
	OrderProduct := OrderProductRepository.Save(context.Background(), tx, domain.OrderProduct{
		Nama: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/OrderProduct/"+strconv.Itoa(OrderProduct.Id), nil)
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

func TestDeleteOrderProductFailed(t *testing.T) {
	db := setupTestDB()
	truncateOrderProduct(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/OrderProduct/404", nil)
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

func TestListOrderProductSucces(t *testing.T) {
	db := setupTestDB()
	truncateOrderProduct(db)

	tx, _ := db.Begin()
	OrderProductRepository := repository.NewOrderProductRepository()
	OrderProduct1 := OrderProductRepository.Save(context.Background(), tx, domain.OrderProduct{
		Nama: "Gadget",
	})
	OrderProduct2:= OrderProductRepository.Save(context.Background(), tx, domain.OrderProduct{
		Nama: "Computer",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/OrderProduct", nil)
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

	OrderProductResponse1 := OrderProduct[0].(map[string]interface{})
	OrderProductResponse2 := OrderProduct[1].(map[string]interface{})

	assert.Equal(t, OrderProduct1.Id, int(OrderProductResponse1["id"].(float64)))
	assert.Equal(t, OrderProduct1.Nama,OrderProductResponse1["name"])

	assert.Equal(t, OrderProduct2.Id, int(OrderProductResponse2["id"].(float64)))
	assert.Equal(t, OrderProduct2.Nama, OrderProductResponse2["name"])
}

func TestUnautorized(t *testing.T) {
	db := setupTestDB()
	truncateOrderProduct(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/OrderProduct", nil)
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

/*
