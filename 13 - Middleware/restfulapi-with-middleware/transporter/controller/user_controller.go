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
		return helper.GetNewResponseFormat(c, http.StatusBadRequest, nil, "")
	}

	user, err := controller.UserService.Create(c.Request().Context(), userRequestBody)
	if err != nil {
		return helper.GetNewResponseFormat(c, http.StatusInternalServerError, nil, "")
	}
	return helper.GetNewResponseFormat(c, http.StatusOK, user, "success, data created")
}

func (controller *UserController) Update(c echo.Context) error {
	userRequestBody := dto.UserRequestBody{}
	err := helper.ReadFromRequestBody(c.Request(), &userRequestBody)
	var id, errId = strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil && errId != nil {
		return helper.GetNewResponseFormat(c, http.StatusBadRequest, nil, "")
	}

	user, err := controller.UserService.Update(c.Request().Context(), userRequestBody, uint(id))
	if err != nil {
		return helper.ResponseChoiser(c, err)
	}
	return helper.GetNewResponseFormat(c, http.StatusOK, user, "success, data updated")
}

func (controller *UserController) Delete(c echo.Context) error {
	var id, errId = strconv.ParseUint(c.Param("id"), 10, 32)
	if errId != nil {
		return helper.GetNewResponseFormat(c, http.StatusBadRequest, nil, "")
	}

	err := controller.UserService.Delete(c.Request().Context(), uint(id))
	if err != nil {
		return helper.ResponseChoiser(c, err)
	}
	return helper.GetNewResponseFormat(c, http.StatusOK, nil, "success, data deleted")
}

func (controller *UserController) FindById(c echo.Context) error {
	var id, err = strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return helper.GetNewResponseFormat(c, http.StatusBadRequest, nil, "")
	}
	user, err := controller.UserService.FindById(c.Request().Context(), uint(id))
	if err != nil {
		return helper.ResponseChoiser(c, err)
	}
	return helper.GetNewResponseFormat(c, http.StatusOK, user, "")
}

func (controller *UserController) FindAll(c echo.Context) error {
	users, err := controller.UserService.FindAll(c.Request().Context())
	if err != nil {
		return helper.ResponseChoiser(c, err)
	}
	return helper.GetNewResponseFormat(c, http.StatusOK, users, "")
}

func (controller *UserController) Auth(c echo.Context) error {
	var authRequest = dto.UserRequestAuth{}
	err := helper.ReadFromRequestBody(c.Request(), &authRequest)
	if err != nil {
		return helper.GetNewResponseFormat(c, http.StatusBadRequest, nil, "")
	}

	user, err := controller.UserService.Auth(c.Request().Context(), authRequest)
	if err != nil {
		return helper.ResponseChoiser(c, err)
	}

	token, err := helper.GenerateToken(user.(dto.UserResponse).Name)
	if err != nil {
		return helper.GetNewResponseFormat(c, http.StatusInternalServerError, nil, "")
	}

	var data = map[string]interface{}{
		"token": token,
		"user":  user,
	}
	return helper.GetNewResponseFormat(c, http.StatusOK, data, "")
}
