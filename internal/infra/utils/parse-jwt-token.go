package utils

import (
	"github.com/golang-jwt/jwt/v5"
	appError "github.com/henrique998/go-auth-2/internal/app/errors"
	"github.com/henrique998/go-auth-2/internal/configs/logger"
)

func ParseJWTToken(cookie, secret string) (string, appError.AppErr) {
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(cookie, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		logger.Error("Error while parse cookie.", err)
		return "", appError.NewAppErr("Internal server error.", 500)
	}

	return claims.Sub, nil
}
