package middleware

import (
	"net/http"
	"strings"

	"github.com/dlyycious/gopos/config"
	"github.com/dlyycious/gopos/internal/module/authmodule"
	"github.com/dlyycious/gopos/pkg/jwtpackage"
	"github.com/dlyycious/gopos/pkg/responsepackage"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(c *fiber.Ctx) error {
	// Ambil header Authorization
	authHeader := c.Get("Authorization")

	// Periksa apakah Authorization header sesuai dengan format "Bearer {token}"
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return responsepackage.SendJSON(c, "unauthorized", http.StatusUnauthorized)
	}

	// Ambil token dari Authorization header
	token := strings.TrimPrefix(authHeader, "Bearer ")
	jwtsplit := strings.Split(token, ".")

	if len(jwtsplit) != 3 {
		return responsepackage.SendJSON(c, "unauthorized", http.StatusUnauthorized)
	}
	verify, err := jwt.ParseWithClaims(token, &jwtpackage.JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.JWT_KEY), nil
	})

	if err != nil {
		return responsepackage.SendJSON(c, "unauthorized", http.StatusUnauthorized)
	}

	if !verify.Valid {
		return responsepackage.SendJSON(c, "unauthorized", http.StatusUnauthorized)
	}

	checktoken := jwtsplit[2]
	user, errors := authmodule.VerifiedToken(checktoken)

	if errors != nil {
		return responsepackage.SendJSON(c, "unauthorized", http.StatusUnauthorized)
	}
	c.Locals("user", user)
	return c.Next()
}
