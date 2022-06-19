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

var urlUser string = "http://localhost:8000/users"

func userSetupTestDB() *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		"root",
		"",
		"localhost",
		"3306",
		"alta_crud_go_test",
	)
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&domain.User{})
	return db
}

func setupRouterUser(db *gorm.DB) *echo.Echo {
	e := echo.New()

	repository := sql.NewUserRepositoryImplSql(db)
	service := service.NewUserServiceImpl(repository)
	controller := controller.NewUserController(service)

	e.GET("/users", controller.FindAll, middleware.IsAuthenticated)
	e.POST("/users", controller.Create)
	e.GET("/users/:id", controller.FindById, middleware.IsAuthenticated)
	e.PUT("/users/:id", controller.Update, middleware.IsAuthenticated)
	e.DELETE("/users/:id", controller.Delete, middleware.IsAuthenticated)
	e.POST("/users/auth", controller.Auth)

	return e
}

func cleanTableUser(db *gorm.DB) {
	db.Exec("TRUNCATE TABLE users")
}

func addUserRequestHeader(r *http.Request, includeJson bool) {
	if includeJson {
		r.Header.Add("Content-Type", "application/json")
	}
	r.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiVGF1ZmlxIiwiZXhwIjoxNjU1ODcyNjg0fQ.38cyvRK3DGOeJcuSXiYdBalY2hFYXOnsGyvnzzelo_Q")
}

func TestCreateUserSuccess(t *testing.T) {
	db := userSetupTestDB()
	cleanTableUser(db)
	router := setupRouterUser(db)

	requestBody := strings.NewReader(`{
		"name" : "baru",
		"email": "baru@example.com",
		"password": "example"
	}`)
	request := httptest.NewRequest(http.MethodPost, urlUser, requestBody)
	addUserRequestHeader(request, true)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "baru", responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, "baru@example.com", responseBody["data"].(map[string]interface{})["email"])
}

func TestCreateUserFailed(t *testing.T) {
	db := userSetupTestDB()
	cleanTableUser(db)
	router := setupRouterUser(db)

	requestBody := strings.NewReader(`{
		"nameeee" : "baru",
		"email": "baru@example.com"		
	}`)
	request := httptest.NewRequest(http.MethodPost, urlUser, requestBody)
	addUserRequestHeader(request, true)

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

func TestUpdateUserSuccess(t *testing.T) {
	db := userSetupTestDB()
	cleanTableUser(db)
	router := setupRouterUser(db)

	// insert baru
	repository := sql.NewUserRepositoryImplSql(db)
	user, err := repository.Save(context.Background(), domain.User{
		Name:     "baru",
		Email:    "baru@example.com",
		Password: "example",
	})

	if err == nil {
		requestBody := strings.NewReader(`{
			"name" : "update",
			"email": "",
			"password": "example"
		}`)
		request := httptest.NewRequest(http.MethodPut, urlUser+"/"+strconv.Itoa(int(user.ID)), requestBody)
		addUserRequestHeader(request, true)

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)
		response := recorder.Result()

		assert.Equal(t, 200, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 200, int(responseBody["code"].(float64)))
		assert.Equal(t, user.ID, uint(responseBody["data"].(map[string]interface{})["id"].(float64)))
		assert.Equal(t, "update", responseBody["data"].(map[string]interface{})["name"])
		assert.Equal(t, user.Email, responseBody["data"].(map[string]interface{})["email"])
	}
}

func TestUpdateUserFailed(t *testing.T) {
	db := userSetupTestDB()
	cleanTableUser(db)
	router := setupRouterUser(db)

	// insert baru
	repository := sql.NewUserRepositoryImplSql(db)
	_, err := repository.Save(context.Background(), domain.User{
		Name:     "baru",
		Email:    "baru@example.com",
		Password: "example",
	})

	if err == nil {
		requestBody := strings.NewReader(`{
			"name" : "update",
			"email": "",
			"password": "example"
		}`)
		request := httptest.NewRequest(http.MethodPut, urlUser+"/2", requestBody)
		addUserRequestHeader(request, true)

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

func TestGetUserSuccess(t *testing.T) {
	db := userSetupTestDB()
	cleanTableUser(db)
	router := setupRouterUser(db)

	// insert baru
	repository := sql.NewUserRepositoryImplSql(db)
	user, err := repository.Save(context.Background(), domain.User{
		Name:     "baru",
		Email:    "baru@example.com",
		Password: "example",
	})

	if err == nil {
		request := httptest.NewRequest(http.MethodGet, urlUser+"/"+strconv.Itoa(int(user.ID)), nil)
		addUserRequestHeader(request, true)

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)
		response := recorder.Result()

		assert.Equal(t, 200, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 200, int(responseBody["code"].(float64)))
		assert.Equal(t, user.ID, uint(responseBody["data"].(map[string]interface{})["id"].(float64)))
		assert.Equal(t, user.Name, responseBody["data"].(map[string]interface{})["name"])
		assert.Equal(t, user.Email, responseBody["data"].(map[string]interface{})["email"])
	}
}

func TestGetUserFailed(t *testing.T) {
	db := userSetupTestDB()
	cleanTableUser(db)
	router := setupRouterUser(db)

	// insert baru
	repository := sql.NewUserRepositoryImplSql(db)
	_, err := repository.Save(context.Background(), domain.User{
		Name:     "baru",
		Email:    "baru@example.com",
		Password: "example",
	})

	if err == nil {
		request := httptest.NewRequest(http.MethodGet, urlUser+"/2", nil)
		addUserRequestHeader(request, true)

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

func TestDeleteUserSuccess(t *testing.T) {
	db := userSetupTestDB()
	cleanTableUser(db)
	router := setupRouterUser(db)

	// insert baru
	repository := sql.NewUserRepositoryImplSql(db)
	user, err := repository.Save(context.Background(), domain.User{
		Name:     "baru",
		Email:    "baru@example.com",
		Password: "example",
	})

	if err == nil {
		request := httptest.NewRequest(http.MethodDelete, urlUser+"/"+strconv.Itoa(int(user.ID)), nil)
		addUserRequestHeader(request, true)

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

func TestDeleteUserFailed(t *testing.T) {
	db := userSetupTestDB()
	cleanTableUser(db)
	router := setupRouterUser(db)

	// insert baru
	repository := sql.NewUserRepositoryImplSql(db)
	_, err := repository.Save(context.Background(), domain.User{
		Name:     "baru",
		Email:    "baru@example.com",
		Password: "example",
	})

	if err == nil {
		request := httptest.NewRequest(http.MethodDelete, urlUser+"/2", nil)
		addUserRequestHeader(request, true)

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

func TestListUserSuccess(t *testing.T) {
	db := userSetupTestDB()
	cleanTableUser(db)
	router := setupRouterUser(db)

	var usersInsert = []domain.User{}

	// insert baru
	repository := sql.NewUserRepositoryImplSql(db)
	user1, err := repository.Save(context.Background(), domain.User{
		Name:     "user1",
		Email:    "user1@example.com",
		Password: "example",
	})
	user2, err2 := repository.Save(context.Background(), domain.User{
		Name:     "user2",
		Email:    "user2@example.com",
		Password: "example",
	})

	if err == nil && err2 == nil {
		usersInsert = append(usersInsert, user1)
		usersInsert = append(usersInsert, user2)

		request := httptest.NewRequest(http.MethodGet, urlUser, nil)
		addUserRequestHeader(request, true)

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)
		response := recorder.Result()

		assert.Equal(t, 200, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 200, int(responseBody["code"].(float64)))
		assert.Equal(t, "success", responseBody["status"])

		var users = responseBody["data"].([]interface{})

		for i := 0; i < len(users); i++ {
			user := users[i].(map[string]interface{})
			assert.Equal(t, usersInsert[i].ID, uint(user["id"].(float64)))
			assert.Equal(t, usersInsert[i].Name, user["name"])
			assert.Equal(t, usersInsert[i].Email, user["email"])
		}

	}
}
