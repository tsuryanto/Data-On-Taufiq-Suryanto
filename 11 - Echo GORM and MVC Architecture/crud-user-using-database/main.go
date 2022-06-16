package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
)

var (
	DB *gorm.DB
)

func init() {
	InitDB()
	InitialMigration()
}

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() {
	config := Config{
		DB_Username: "root",
		DB_Password: "",
		DB_Port:     "3306",
		DB_Host:     "localhost",
		DB_Name:     "alta_crud_go",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
}

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func InitialMigration() {
	DB.AutoMigrate(&User{})
}

// get all users
func GetUsersController(c echo.Context) error {
	var users []User

	if err := DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	var user = User{}
	var id, err = strconv.Atoi(c.Param("id"))
	if err == nil {
		if err := DB.First(&user, id).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get user information",
			"user":    user,
		})
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
}

// create new user
func CreateUserController(c echo.Context) error {
	var user = User{}
	c.Bind(&user)

	if err := DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"user":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	var id, err = strconv.Atoi(c.Param("id"))
	if err == nil {
		if err := DB.Delete(&User{}, id).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success user deleted",
		})
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
}

// delete user by id
func UpdateUserController(c echo.Context) error {
	var user = User{}
	c.Bind(&user)

	var id, err = strconv.Atoi(c.Param("id"))
	if err == nil {
		if err := DB.Model(&user).Where("id = ?", id).Updates(&user).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := DB.First(&user, id).Error; err != nil {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "success update user information",
				"user":    "sorry, can't show this value",
			})
		} else {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "success update user information",
				"user":    user,
			})
		}
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
}

func main() {
	e := echo.New()

	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)
	e.GET("/users/:id", GetUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.DELETE("/users/:id", DeleteUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
