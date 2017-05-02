package tokens

import (
	"math/rand"
	"time"
)

var (
	alphaChars        = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	alphaNumericChars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

func RandomAlphaString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = alphaChars[rand.Intn(len(alphaChars))]
	}
	return string(b)
}

func RandomAlphaNumericString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = alphaNumericChars[rand.Intn(len(alphaNumericChars))]
	}
	return string(b)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
