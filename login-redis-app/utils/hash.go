package utils

import (
	"crypto/sha1"
	"fmt"
)

// HashSHA1 mengubah string menjadi hash SHA1
func HashSHA1(text string) string {
	h := sha1.New()
	h.Write([]byte(text))
	return fmt.Sprintf("%x", h.Sum(nil))
}
