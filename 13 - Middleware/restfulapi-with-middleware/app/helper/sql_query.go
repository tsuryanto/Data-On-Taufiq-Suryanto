package helper

import (
	"errors"
	"net/http"

	"gorm.io/gorm"
)

func GetQueryErrorResponseCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusNotFound
	}
	return http.StatusInternalServerError
}
