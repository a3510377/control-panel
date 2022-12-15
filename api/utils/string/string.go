package string

import (
	"math/rand"
	"time"
)

const (
	Whitespace     = " \t\n\r\v\f"
	AsciiLowercase = "abcdefghijklmnopqrstuvwxyz"
	AsciiUppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	AsciiLetters   = AsciiLowercase + AsciiUppercase
	Digits         = "0123456789"
	HexDigits      = Digits + "abcdef" + "ABCDEF"
	OctDigits      = "01234567"
	Punctuation    = `!"#$%&'()*+,-./:;<=>?@[\]^_{|}~` + "`"
	Printable      = Digits + AsciiLowercase + Punctuation + Whitespace
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomString(n int) string {
	r := ""
	for i := 0; i < n; i++ {
		r += string(Printable[rand.Intn(len(Printable)-1)])
	}
	return r
}
