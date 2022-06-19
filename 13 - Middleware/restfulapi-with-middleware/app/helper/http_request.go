package helper

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ReadFromRequestBody(request *http.Request, result interface{}) error {
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(result); err != nil {
		return err
	} else {
		validate := validator.New()
		err := validate.Struct(result)
		// fmt.Println(err)
		if err != nil {
			return err
		} else {
			return nil
		}
	}
}
