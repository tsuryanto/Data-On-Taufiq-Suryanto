package constant

import "errors"

var ErrNotFound error = errors.New("not found")
var ErrBadRequest error = errors.New("bad request")
var ErrServerError error = errors.New("server error")
var ErrUnauthorized error = errors.New("unauthorized")
