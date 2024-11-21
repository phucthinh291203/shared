package token

import (
	"time"

	errors "github.com/phucthinh291203/shared/errors"

	"github.com/dgrijalva/jwt-go"
)

type BaseClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Name     string `json:"name"`
	UserID   string `json:"user_id"`
	jwt.StandardClaims
}

func GenerateJWT(claims BaseClaims, secretKey string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims.ExpiresAt = expirationTime.Unix() // Thiết lập thời gian hết hạn

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))

}

func ParseJWT(tokenString string, secretKey string) (*BaseClaims, error) {
	claims := &BaseClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.ErrorWithMessage(errors.ErrInternalServerError, "Token quá hạn")
	}
	return claims, nil
}
