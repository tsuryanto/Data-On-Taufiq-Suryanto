package controller

import (
	"latian_clean_arch/app/helper"
	"latian_clean_arch/model/dto"
	"latian_clean_arch/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookController struct {
	BookService service.BookService
}

func NewBookController(service service.BookService) *BookController {
	return &BookController{
		BookService: service,
	}
}

func (controller *BookController) Create(c echo.Context) error {
	bookRequestBody := dto.BookRequestBody{}
	err := helper.ReadFromRequestBody(c.Request(), &bookRequestBody)
	if err != nil {
		var code = http.StatusBadRequest
		return c.JSON(code, helper.GetResponseFormat(code, nil))
	}

	code, book, _ := controller.BookService.Create(c.Request().Context(), bookRequestBody)
	return c.JSON(code, helper.GetResponseFormat(code, book))
}

func (controller *BookController) Update(c echo.Context) error {
	bookRequestBody := dto.BookRequestBody{}
	err := helper.ReadFromRequestBody(c.Request(), &bookRequestBody)
	var id, errId = strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil && errId != nil {
		var code = http.StatusBadRequest
		return c.JSON(code, helper.GetResponseFormat(code, nil))
	}

	code, book, _ := controller.BookService.Update(c.Request().Context(), bookRequestBody, uint(id))
	return c.JSON(code, helper.GetResponseFormat(code, book))
}

func (controller *BookController) Delete(c echo.Context) error {
	var id, errId = strconv.ParseUint(c.Param("id"), 10, 32)
	if errId != nil {
		var code = http.StatusBadRequest
		return c.JSON(code, helper.GetResponseFormatForDelete(code))
	}

	code, _ := controller.BookService.Delete(c.Request().Context(), uint(id))
	return c.JSON(code, helper.GetResponseFormatForDelete(code))
}

func (controller *BookController) FindById(c echo.Context) error {
	var id, err = strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		var code = http.StatusBadRequest
		return c.JSON(code, helper.GetResponseFormat(code, nil))
	}
	code, book, _ := controller.BookService.FindById(c.Request().Context(), uint(id))
	return c.JSON(code, helper.GetResponseFormat(code, book))
}

func (controller *BookController) FindAll(c echo.Context) error {
	code, books, _ := controller.BookService.FindAll(c.Request().Context())
	return c.JSON(code, helper.GetResponseFormat(code, books))
}
