package controller

import (
	"latian_clean_arch/app/helper"
	"latian_clean_arch/model/dto"
	"latian_clean_arch/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{
		UserService: service,
	}
}

func (controller *UserController) Create(c echo.Context) error {
	userRequestBody := dto.UserRequestBody{}
	err := helper.ReadFromRequestBody(c.Request(), &userRequestBody)
	if err != nil {
		var code = http.StatusBadRequest
		return c.JSON(code, helper.GetResponseFormat(code, nil))
	}

	code, user, _ := controller.UserService.Create(c.Request().Context(), userRequestBody)
	return c.JSON(code, helper.GetResponseFormat(code, user))
}

func (controller *UserController) Update(c echo.Context) error {
	userRequestBody := dto.UserRequestBody{}
	err := helper.ReadFromRequestBody(c.Request(), &userRequestBody)
	var id, errId = strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil && errId != nil {
		var code = http.StatusBadRequest
		return c.JSON(code, helper.GetResponseFormat(code, nil))
	}

	code, user, _ := controller.UserService.Update(c.Request().Context(), userRequestBody, uint(id))
	return c.JSON(code, helper.GetResponseFormat(code, user))
}

func (controller *UserController) Delete(c echo.Context) error {
	var id, errId = strconv.ParseUint(c.Param("id"), 10, 32)
	if errId != nil {
		var code = http.StatusBadRequest
		return c.JSON(code, helper.GetResponseFormatForDelete(code))
	}

	code, _ := controller.UserService.Delete(c.Request().Context(), uint(id))
	return c.JSON(code, helper.GetResponseFormatForDelete(code))
}

func (controller *UserController) FindById(c echo.Context) error {
	var id, err = strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		var code = http.StatusBadRequest
		return c.JSON(code, helper.GetResponseFormat(code, nil))
	}
	code, user, _ := controller.UserService.FindById(c.Request().Context(), uint(id))
	return c.JSON(code, helper.GetResponseFormat(code, user))
}

func (controller *UserController) FindAll(c echo.Context) error {
	code, users, _ := controller.UserService.FindAll(c.Request().Context())
	return c.JSON(code, helper.GetResponseFormat(code, users))
}

func (controller *UserController) Auth(c echo.Context) error {
	var authRequest = dto.UserRequestAuth{}
	err := helper.ReadFromRequestBody(c.Request(), &authRequest)
	if err != nil {
		var code = http.StatusBadRequest
		return c.JSON(code, helper.GetResponseFormatForDelete(code))
	}

	code, user, _ := controller.UserService.Auth(c.Request().Context(), authRequest)
	if code == http.StatusOK {
		token, err := helper.GenerateToken(user.(dto.UserResponse).Name)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.GetResponseFormat(http.StatusInternalServerError, nil))
		}

		var data = map[string]interface{}{
			"token": token,
			"user":  user,
		}
		return c.JSON(code, helper.GetResponseFormat(code, data))
	} else {
		return c.JSON(code, helper.GetResponseFormat(code, nil))
	}
}
