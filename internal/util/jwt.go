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

func ValidateToken(token string) error {
	tk, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SIGNING_KEY")), nil
	})

	if err != nil || !tk.Valid {
		return err
	} else if _, ok := tk.Claims.(*tokenClaims); ok {
		if expTime, err := tk.Claims.GetExpirationTime(); err != nil || expTime.Time.Compare(time.Now()) == -1 {
			return ErrorExpiredToken
		}

		return nil
	} else {
		return ErrorUnknownClaimsType
	}
}
