package cfg

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"time"
)

func JWTAuthMiddleware(c *fiber.Ctx) error {
	// Ekstrak token dari header
	token := extractToken(c)
	if token == "" {
		return unauthorized(c, "No token provided")
	}

	// Parse dan validasi token
	claims, err := parseToken(token)
	if err != nil {
		return unauthorized(c, err.Error())
	}

	// Tambahan validasi tambahan
	if err := validateTokenClaims(claims); err != nil {
		return unauthorized(c, err.Error())
	}

	// Set claims di context untuk digunakan di handler selanjutnya
	c.Locals("username", claims.Username)
	c.Locals("claims", claims)

	return c.Next()
}

func extractToken(c *fiber.Ctx) string {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ""
	}

	return parts[1]
}

func parseToken(tokenString string) (*JWTClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_KEY), nil
	})

	if err != nil {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}

func validateTokenClaims(claims *JWTClaim) error {
	// Validasi waktu kedaluwarsa
	if claims.ExpiresAt.Time.Before(time.Now()) {
		return errors.New("token has expired")
	}

	// Validasi issuer (opsional)
	if claims.Issuer != "koriebruh.akaJamal" {
		return errors.New("invalid token issuer")
	}

	return nil
}

func unauthorized(c *fiber.Ctx, message string) error {
	return c.Status(http.StatusUnauthorized).JSON(
		struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}{Status: "UNAUTHORIZED", Message: message},
	)
}
