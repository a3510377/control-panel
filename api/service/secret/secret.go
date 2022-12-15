package secret

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWT string

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
type RefreshToken struct {
	Token          string    `json:"token"`
	ExpirationTime time.Time `json:"expiration"`
}

var jwtKey = os.Getenv("jwt_secret")

func NewClaims() *Claims { return &Claims{} }

func Create(username string, time time.Duration) {
}

// JWT to string
func (j JWT) String() string { return string(j) }

// state: `200` OK
// state: `400` Bad Request
// state: `401` Unauthorized
// data: token info
func (j JWT) Info() (data *Claims, status int) {
	claims := NewClaims()
	token, err := jwt.ParseWithClaims(j.String(), claims, func(token *jwt.Token) (any, error) {
		return jwtKey, nil
	})

	if err == jwt.ErrSignatureInvalid || !token.Valid {
		return nil, http.StatusUnauthorized
	} else if err != nil {
		return nil, http.StatusBadRequest
	}

	return claims, http.StatusOK
}

func (j JWT) Refresh(oldToken string, newTime time.Duration) (token *RefreshToken, status int) {
	claims, status := j.Info()
	if status != 200 {
		return nil, status
	}

	expirationTime := time.Now().Add(newTime)
	claims.RegisteredClaims = jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(expirationTime)}

	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtKey)
	if err != nil {
		return nil, 500
	}

	return &RefreshToken{tokenString, expirationTime}, status
}
