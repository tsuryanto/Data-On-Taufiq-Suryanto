package middleware

import (
	. "latian_clean_arch/app/constant"
	"latian_clean_arch/app/helper"

	"github.com/labstack/echo/v4/middleware"
)

var config = middleware.JWTConfig{
	Claims:     &helper.JwtCustomClaims{},
	SigningKey: []byte(SECRET_JWT),
}
var IsAuthenticated = middleware.JWTWithConfig(config)
