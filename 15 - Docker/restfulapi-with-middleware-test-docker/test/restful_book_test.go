package test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"latian_clean_arch/app/middleware"
	"latian_clean_arch/model/domain"
	"latian_clean_arch/repository/sql"
	"latian_clean_arch/service"
	"latian_clean_arch/transporter/controller"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var urlBook string = "http://localhost:8000/books"

func bookSetupTestDB() *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		"root",
		"",
		"host.docker.internal",
		"3306",
		"alta_crud_go_test",
	)
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&domain.Book{})
	return db
}

func setupRouterBook(db *gorm.DB) *echo.Echo {
	e := echo.New()

	repository := sql.NewBookRepositoryImplSql(db)
	service := service.NewBookServiceImpl(repository)
	controller := controller.NewBookController(service)

	e.GET("/books", controller.FindAll, middleware.IsAuthenticated)
	e.POST("/books", controller.Create)
	e.GET("/books/:id", controller.FindById, middleware.IsAuthenticated)
	e.PUT("/books/:id", controller.Update, middleware.IsAuthenticated)
	e.DELETE("/books/:id", controller.Delete, middleware.IsAuthenticated)

	return e
}

func cleanTableBook(db *gorm.DB) {
	db.Exec("TRUNCATE TABLE books")
}

func addBookRequestHeader(r *http.Request, includeJson bool) {
	if includeJson {
		r.Header.Add("Content-Type", "application/json")
	}
	r.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiVGF1ZmlxIiwiZXhwIjoxNjU1ODcyNjg0fQ.38cyvRK3DGOeJcuSXiYdBalY2hFYXOnsGyvnzzelo_Q")
}

func TestCreateBookSuccess(t *testing.T) {
	db := bookSetupTestDB()
	cleanTableBook(db)
	router := setupRouterBook(db)

	requestBody := strings.NewReader(`{
		"name" : "Untuk Kamu",
		"publisher": "Gramedia",
		"author": "Tere Liye"
	}`)
	request := httptest.NewRequest(http.MethodPost, urlBook, requestBody)
	addBookRequestHeader(request, true)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "Untuk Kamu", responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, "Gramedia", responseBody["data"].(map[string]interface{})["publisher"])
	assert.Equal(t, "Tere Liye", responseBody["data"].(map[string]interface{})["author"])
}

func TestCreateBookFailed(t *testing.T) {
	db := bookSetupTestDB()
	cleanTableBook(db)
	router := setupRouterBook(db)

	requestBody := strings.NewReader(`{
		"nameeee" : "Untuk Kamu",
		"publisher": "Gramedia"		
	}`)
	request := httptest.NewRequest(http.MethodPost, urlBook, requestBody)
	addBookRequestHeader(request, true)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 400, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "bad request", responseBody["message"])
}

func TestUpdateBookSuccess(t *testing.T) {
	db := bookSetupTestDB()
	cleanTableBook(db)
	router := setupRouterBook(db)

	// insert baru
	repository := sql.NewBookRepositoryImplSql(db)
	book, err := repository.Save(context.Background(), domain.Book{
		Name:      "Untuk Kamu",
		Publisher: "Gramedia",
		Author:    "Tere Liye",
	})

	if err == nil {
		requestBody := strings.NewReader(`{
			"name" : "update",
			"author": "Budiman"
		}`)
		request := httptest.NewRequest(http.MethodPut, urlBook+"/"+strconv.Itoa(int(book.ID)), requestBody)
		addBookRequestHeader(request, true)

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)
		response := recorder.Result()

		assert.Equal(t, 200, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 200, int(responseBody["code"].(float64)))
		assert.Equal(t, book.ID, uint(responseBody["data"].(map[string]interface{})["id"].(float64)))
		assert.Equal(t, "update", responseBody["data"].(map[string]interface{})["name"])
		assert.Equal(t, book.Publisher, responseBody["data"].(map[string]interface{})["publisher"])
		assert.Equal(t, "Budiman", responseBody["data"].(map[string]interface{})["author"])
	}
}

func TestUpdateBookFailed(t *testing.T) {
	db := bookSetupTestDB()
	cleanTableBook(db)
	router := setupRouterBook(db)

	// insert baru
	repository := sql.NewBookRepositoryImplSql(db)
	_, err := repository.Save(context.Background(), domain.Book{
		Name:      "Untuk Kamu",
		Publisher: "Gramedia",
		Author:    "Tere Liye",
	})

	if err == nil {
		requestBody := strings.NewReader(`{
			"name" : "update",
			"author": "Budiman"
		}`)
		request := httptest.NewRequest(http.MethodPut, urlBook+"/2", requestBody)
		addBookRequestHeader(request, true)

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)
		response := recorder.Result()

		assert.Equal(t, 404, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 404, int(responseBody["code"].(float64)))
		assert.Equal(t, "not found", responseBody["status"])
		assert.Nil(t, responseBody["data"])
	}
}

func TestGetBookSuccess(t *testing.T) {
	db := bookSetupTestDB()
	cleanTableBook(db)
	router := setupRouterBook(db)

	// insert baru
	repository := sql.NewBookRepositoryImplSql(db)
	book, err := repository.Save(context.Background(), domain.Book{
		Name:      "Untuk Kamu",
		Publisher: "Gramedia",
		Author:    "Tere Liye",
	})

	if err == nil {
		request := httptest.NewRequest(http.MethodGet, urlBook+"/"+strconv.Itoa(int(book.ID)), nil)
		addBookRequestHeader(request, true)

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)
		response := recorder.Result()

		assert.Equal(t, 200, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 200, int(responseBody["code"].(float64)))
		assert.Equal(t, book.ID, uint(responseBody["data"].(map[string]interface{})["id"].(float64)))
		assert.Equal(t, book.Name, responseBody["data"].(map[string]interface{})["name"])
		assert.Equal(t, book.Publisher, responseBody["data"].(map[string]interface{})["publisher"])
		assert.Equal(t, book.Author, responseBody["data"].(map[string]interface{})["author"])
	}
}

func TestGetBookFailed(t *testing.T) {
	db := bookSetupTestDB()
	cleanTableBook(db)
	router := setupRouterBook(db)

	// insert baru
	repository := sql.NewBookRepositoryImplSql(db)
	_, err := repository.Save(context.Background(), domain.Book{
		Name:      "Untuk Kamu",
		Publisher: "Gramedia",
		Author:    "Tere Liye",
	})

	if err == nil {
		request := httptest.NewRequest(http.MethodGet, urlBook+"/2", nil)
		addBookRequestHeader(request, true)

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)
		response := recorder.Result()

		assert.Equal(t, 404, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 404, int(responseBody["code"].(float64)))
		assert.Equal(t, "not found", responseBody["status"])
		assert.Nil(t, responseBody["data"])
	}
}

func TestDeleteBookSuccess(t *testing.T) {
	db := bookSetupTestDB()
	cleanTableBook(db)
	router := setupRouterBook(db)

	// insert baru
	repository := sql.NewBookRepositoryImplSql(db)
	book, err := repository.Save(context.Background(), domain.Book{
		Name:      "Untuk Kamu",
		Publisher: "Gramedia",
		Author:    "Tere Liye",
	})

	if err == nil {
		request := httptest.NewRequest(http.MethodDelete, urlBook+"/"+strconv.Itoa(int(book.ID)), nil)
		addBookRequestHeader(request, true)

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)
		response := recorder.Result()

		assert.Equal(t, 200, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 200, int(responseBody["code"].(float64)))
		assert.Equal(t, "success", responseBody["status"])
	}
}

func TestDeleteBookFailed(t *testing.T) {
	db := bookSetupTestDB()
	cleanTableBook(db)
	router := setupRouterBook(db)

	// insert baru
	repository := sql.NewBookRepositoryImplSql(db)
	_, err := repository.Save(context.Background(), domain.Book{
		Name:      "Untuk Kamu",
		Publisher: "Gramedia",
		Author:    "Tere Liye",
	})

	if err == nil {
		request := httptest.NewRequest(http.MethodDelete, urlBook+"/2", nil)
		addBookRequestHeader(request, true)

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)
		response := recorder.Result()

		assert.Equal(t, 404, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 404, int(responseBody["code"].(float64)))
		assert.Equal(t, "not found", responseBody["status"])
		assert.Nil(t, responseBody["data"])
	}
}

func TestListBookSuccess(t *testing.T) {
	db := bookSetupTestDB()
	cleanTableBook(db)
	router := setupRouterBook(db)

	var booksInsert = []domain.Book{}

	// insert baru
	repository := sql.NewBookRepositoryImplSql(db)
	book1, err := repository.Save(context.Background(), domain.Book{
		Name:      "book1",
		Publisher: "book1@example.com",
		Author:    "Tere Liye",
	})
	book2, err2 := repository.Save(context.Background(), domain.Book{
		Name:      "book2",
		Publisher: "book2@example.com",
		Author:    "Tere Liye",
	})

	if err == nil && err2 == nil {
		booksInsert = append(booksInsert, book1)
		booksInsert = append(booksInsert, book2)

		request := httptest.NewRequest(http.MethodGet, urlBook, nil)
		addBookRequestHeader(request, true)

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)
		response := recorder.Result()

		assert.Equal(t, 200, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 200, int(responseBody["code"].(float64)))
		assert.Equal(t, "success", responseBody["status"])

		var books = responseBody["data"].([]interface{})

		for i := 0; i < len(books); i++ {
			book := books[i].(map[string]interface{})
			assert.Equal(t, booksInsert[i].ID, uint(book["id"].(float64)))
			assert.Equal(t, booksInsert[i].Name, book["name"])
			assert.Equal(t, booksInsert[i].Publisher, book["publisher"])
			assert.Equal(t, booksInsert[i].Author, book["author"])
		}

	}
}
