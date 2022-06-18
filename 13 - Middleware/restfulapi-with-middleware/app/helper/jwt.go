package helper

import (
	"latian_clean_arch/app/constant"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtCustomClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func GenerateToken(name string) (string, error) {
	claims := &JwtCustomClaims{
		name,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(constant.SECRET_JWT))
	if err != nil {
		return "", err
	}
	return t, nil
}
