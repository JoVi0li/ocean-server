package util

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserInfoToken struct {
	ID       string
	Username string
	Email    string
}

type tokenClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

func NewToken(infos UserInfoToken) (string, error) {
	claims := tokenClaims{
		infos.ID,
		infos.Username,
		infos.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour * 7)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    os.Getenv("JWT_ISSUER"),
			Subject:   os.Getenv("JWT_SUBJECT"),
			ID:        infos.ID,
			Audience:  []string{os.Getenv("JWT_AUDIENCE")},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(os.Getenv("JWT_SIGNING_KEY")))

	if err != nil {
		return "", err
	}

	return ss, nil
}
