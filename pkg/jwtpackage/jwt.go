package jwtpackage

import (
	"time"

	"github.com/dlyycious/gopos/config"
	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	Username string
	Goid     uint
	jwt.RegisteredClaims
}

var expires = time.Now().Add(config.JWT_EXP * time.Hour)

func (j *JWTClaims) Generate() (string, error) {

	claims := JWTClaims{
		j.Username,
		j.Goid,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expires),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "GoPOS",
		},
	}

	createToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := createToken.SignedString([]byte(config.JWT_KEY))

	return token, err
}
