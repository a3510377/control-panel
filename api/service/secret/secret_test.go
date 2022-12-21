package secret

import (
	"testing"
	"time"
)

func TestJWT(t *testing.T) {
	t.Setenv("jwt_secret", "abcdefghijklmnopqrstuvwxyzABCDEF")

	token, _ := New(time.Hour * 1)
	data, _ := token.Token.Info()

	t.Log(token.Token)
	t.Log(data)
	t.Log(token.Token.Refresh(time.Hour * 10))
}
