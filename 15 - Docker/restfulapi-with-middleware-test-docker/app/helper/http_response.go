package helper

import (
	. "latian_clean_arch/app/constant"
	"latian_clean_arch/model/basic"
	"net/http"

	"github.com/labstack/echo/v4"
)

// type HTTPResponseFormat interface {
// 	GetResponse(data interface{}) basic.ResponseFormat
// }

// type HTTPResponseOK {

// }

func ResponseChoiser(c echo.Context, err error) error {
	if err == ErrNotFound {
		return GetNewResponseFormat(c, http.StatusNotFound, nil, "")
	}

	if err == ErrBadRequest {
		return GetNewResponseFormat(c, http.StatusBadRequest, nil, "")
	}

	if err == ErrUnauthorized {
		return GetNewResponseFormat(c, http.StatusUnauthorized, nil, "")
	}

	return GetNewResponseFormat(c, http.StatusInternalServerError, nil, "")
}

func GetNewResponseFormat(c echo.Context, code int, data interface{}, message string) error {
	var responseFormat = basic.ResponseFormat{
		Code:    code,
		Message: message,
		Data:    data,
	}

	switch code {
	case http.StatusOK:
		responseFormat.Status = "success"
		if message == "" {
			responseFormat.Message = "success"
		}

	case http.StatusBadRequest:
		responseFormat.Status = "failed"
		if message == "" {
			responseFormat.Message = "bad request"
		}

	case http.StatusNotFound:
		responseFormat.Status = "not found"
		if message == "" {
			responseFormat.Message = "record not found"
		}

	case http.StatusUnauthorized:
		responseFormat.Status = "failed"
		if message == "" {
			responseFormat.Message = "unauthorized"
		}

	default:
		responseFormat.Status = "failed"
		if message == "" {
			responseFormat.Message = "internal server error"
		}
	}

	return c.JSON(code, responseFormat)
}
