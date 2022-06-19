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
		return helper.GetNewResponseFormat(c, http.StatusBadRequest, nil, "")
	}

	book, err := controller.BookService.Create(c.Request().Context(), bookRequestBody)
	if err != nil {
		return helper.GetNewResponseFormat(c, http.StatusInternalServerError, nil, "")
	}
	return helper.GetNewResponseFormat(c, http.StatusOK, book, "success, data created")
}

func (controller *BookController) Update(c echo.Context) error {
	bookRequestBody := dto.UpdateBookRequestBody{}
	err := helper.ReadFromRequestBody(c.Request(), &bookRequestBody)
	var id, errId = strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil || errId != nil {
		return helper.GetNewResponseFormat(c, http.StatusBadRequest, nil, "")
	}

	book, err := controller.BookService.Update(c.Request().Context(), bookRequestBody, uint(id))
	if err != nil {
		return helper.ResponseChoiser(c, err)
	}
	return helper.GetNewResponseFormat(c, http.StatusOK, book, "success, data updated")
}

func (controller *BookController) Delete(c echo.Context) error {
	var id, errId = strconv.ParseUint(c.Param("id"), 10, 32)
	if errId != nil {
		return helper.GetNewResponseFormat(c, http.StatusBadRequest, nil, "")
	}

	err := controller.BookService.Delete(c.Request().Context(), uint(id))
	if err != nil {
		return helper.ResponseChoiser(c, err)
	}
	return helper.GetNewResponseFormat(c, http.StatusOK, nil, "success, data deleted")
}

func (controller *BookController) FindById(c echo.Context) error {
	var id, err = strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return helper.GetNewResponseFormat(c, http.StatusBadRequest, nil, "")
	}
	book, err := controller.BookService.FindById(c.Request().Context(), uint(id))
	if err != nil {
		return helper.ResponseChoiser(c, err)
	}
	return helper.GetNewResponseFormat(c, http.StatusOK, book, "")
}

func (controller *BookController) FindAll(c echo.Context) error {
	books, err := controller.BookService.FindAll(c.Request().Context())
	if err != nil {
		return helper.ResponseChoiser(c, err)
	}
	return helper.GetNewResponseFormat(c, http.StatusOK, books, "")
}
