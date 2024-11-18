package token

import (
	"time"
	"user-service/config"
	"user-service/errors"

	"github.com/dgrijalva/jwt-go"
)

var SecretKey = config.GetConfig().SecretKey

type BaseClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Name     string `json:"name"`
	UserID   string `json:"user_id"`
	jwt.StandardClaims
}

func GenerateJWT(claims BaseClaims) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims.ExpiresAt = expirationTime.Unix() // Thiết lập thời gian hết hạn

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SecretKey)

}

func ParseJWT(tokenString string) (*BaseClaims, error) {
	claims := &BaseClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.ErrorWithMessage(errors.ErrInternalServerError, "Token quá hạn")
	}
	return claims, nil
}
