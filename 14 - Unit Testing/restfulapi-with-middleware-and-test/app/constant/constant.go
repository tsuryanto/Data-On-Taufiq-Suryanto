package constant

import (
	"encoding/base64"
	"errors"
)

var SECRET_JWT string = base64.StdEncoding.EncodeToString([]byte("golang"))
var ErrNotFound error = errors.New("not found")
var ErrBadRequest error = errors.New("bad request")
var ErrServerError error = errors.New("server error")
var ErrUnauthorized error = errors.New("unauthorized")
