package secret

import (
	"testing"
	"time"

	"github.com/a3510377/control-panel/utils/string"
)

func TestJWT(t *testing.T) {
	t.Setenv("jwt_secret", string.Printable)

	token, _ := New(time.Hour * 1)
	data, _ := token.Token.Info()

	t.Log(token.Token)
	t.Log(data)
	t.Log(token.Token.Refresh(time.Hour * 10))
}
