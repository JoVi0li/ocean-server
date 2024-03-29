package shared

import (
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserInfoToken struct {
	ID       string
	Username string
	Email    string
}

type TokenClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

func NewToken(infos UserInfoToken) (string, error) {
	claims := TokenClaims{
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
	tk, err := jwt.ParseWithClaims(token, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SIGNING_KEY")), nil
	})

	if err != nil || !tk.Valid {
		return err
	} else if _, ok := tk.Claims.(*TokenClaims); ok {
		if expTime, err := tk.Claims.GetExpirationTime(); err != nil || expTime.Time.Compare(time.Now()) == -1 {
			return ErrorExpiredToken
		}
		return nil
	} else {
		return ErrorUnknownClaimsType
	}
}

func GetToken(ctx *gin.Context) (string, error) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		return "", ErrorMissingAuthorizationToken
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
	if tokenString == "" {
		return "", ErrorInvalidAuthorizationToken
	}

	return tokenString, nil
}

func DecodeTokenClaims(token string) (*TokenClaims, error) {
	tk, err := jwt.ParseWithClaims(token, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SIGNING_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	tkClaims, ok := tk.Claims.(*TokenClaims)
	if !ok {
		return nil, ErrorUnknownClaimsType
	}

	return tkClaims, nil
}
