package utils

import (
	"crypto/rand"
	"math/big"
)

var alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateShortId() string {
	key := make([]rune, 16)

	mx := big.NewInt(int64(len(alphabet)))

	for i := 0; i < len(key); i++ {
		b, _ := rand.Int(rand.Reader, mx)
		key[i] = rune(alphabet[b.Int64()])
	}

	return string(key)
}
