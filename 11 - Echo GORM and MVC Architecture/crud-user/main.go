package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// ------------------------ Controller -------------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		for i := range users {
			if users[i].Id == id {
				return c.JSON(http.StatusOK, map[string]interface{}{
					"messages": "success get user",
					"user":     users[i],
				})
			}
		}
	} else {
		return err
	}
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"messages": "not found",
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	var id, err = strconv.Atoi(c.Param("id"))
	if err == nil {
		for i := range users {
			if users[i].Id == id {
				users = append(users[:i], users[i+1:]...)
				return c.JSON(http.StatusOK, map[string]interface{}{
					"messages": "success, user deleted",
					"user":     users,
				})
			}
		}
	} else {
		return err
	}
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"messages": "not found",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	var id, err = strconv.Atoi(c.Param("id"))
	var user = User{}
	c.Bind(&user)

	if err == nil {
		for i := range users {
			if users[i].Id == id {
				if user.Name != "" {
					users[i].Name = user.Name
				}
				if user.Email != "" {
					users[i].Name = user.Email
				}
				if user.Password != "" {
					users[i].Name = user.Password
				}
				return c.JSON(http.StatusOK, map[string]interface{}{
					"messages": "success update user information",
					"user":     users[i],
				})
			}
		}
	} else {
		return err
	}
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"messages": "not found",
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create users",
		"users":    users,
	})
}

// ---------------------------------------------------------

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
