package password

import (
	"crypto/rand"
	"crypto/sha256"
)

var (
	defaul = &Password{
		NewHash:    sha256.New,
		RandReader: rand.Reader,
		HashSize:   sha256.Size,
		RandSize:   8,
	}
)

// Encrypt encryption adds random variables to make each encryption different
func Encrypt(word string) (code string) {
	return defaul.Encrypt(word)
}

// Verify password
func Verify(word, code string) bool {
	return defaul.Verify(word, code)
}
