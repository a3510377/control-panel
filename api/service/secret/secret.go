package secret

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte(os.Getenv("jwt_secret"))

type (
	JWT    string
	Claims struct {
		Username string `json:"username"`
		jwt.RegisteredClaims
	}
	RefreshToken struct {
		Token          JWT       `json:"token"`
		ExpirationTime time.Time `json:"expiration"`
	}
)

func New(newTime time.Duration) (token *RefreshToken, status int) { return Create(&Claims{}, newTime) }

func Create(claims *Claims, newTime time.Duration) (token *RefreshToken, status int) {
	expirationTime := time.Now().Add(newTime)

	claims.RegisteredClaims = jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(expirationTime)}

	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtKey)
	if err != nil {
		return nil, 500
	}

	return &RefreshToken{JWT(tokenString), expirationTime}, 200
}

// JWT to string
func (j JWT) String() string { return string(j) }

// state: `200` OK
// state: `400` Bad Request
// state: `401` Unauthorized
// data: token info
func (j JWT) Info() (data *Claims, status int) {
	claims := &Claims{}
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

// Refresh Token
func (j *JWT) Refresh(newTime time.Duration) (refreshToken *RefreshToken, status int) {
	claims, status := j.Info()
	if status != 200 {
		return nil, status
	}

	token, status := Create(claims, newTime)
	if status != http.StatusOK {
		return nil, status
	}

	*j = JWT(token.Token)

	return token, status
}
