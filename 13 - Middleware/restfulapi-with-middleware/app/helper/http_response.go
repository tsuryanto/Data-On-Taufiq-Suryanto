package helper

import (
	"latian_clean_arch/model/basic"
	"net/http"
)

// type HTTPResponseFormat interface {
// 	GetResponse(data interface{}) basic.ResponseFormat
// }

// type HTTPResponseOK {

// }

func GetResponseFormat(code int, data interface{}) basic.ResponseFormat {
	var responseFormat = basic.ResponseFormat{
		Code: code,
		Data: data,
	}

	switch code {
	case http.StatusOK:
		responseFormat.Status = "success"
		responseFormat.Message = "success"

	case http.StatusBadRequest:
		responseFormat.Status = "failed"
		responseFormat.Message = "bad request"

	case http.StatusNotFound:
		responseFormat.Status = "not found"
		responseFormat.Message = "record not found"

	case http.StatusUnauthorized:
		responseFormat.Status = "failed"
		responseFormat.Message = "Unauthorized"

	default:
		responseFormat.Status = "failed"
		responseFormat.Message = "internal server error"
	}

	return responseFormat
}

func GetResponseFormatForDelete(code int) basic.ResponseFormat {
	var responseFormat = basic.ResponseFormat{
		Code: code,
	}

	switch code {
	case http.StatusOK:
		responseFormat.Status = "success"
		responseFormat.Message = "success, data deleted"

	case http.StatusBadRequest:
		responseFormat.Status = "failed"
		responseFormat.Message = "bad request"

	case http.StatusNotFound:
		responseFormat.Status = "not found"
		responseFormat.Message = "record not found"

	case http.StatusUnauthorized:
		responseFormat.Status = "failed"
		responseFormat.Message = "Unauthorized"

	default:
		responseFormat.Status = "failed"
		responseFormat.Message = "internal server error"
	}

	return responseFormat
}
