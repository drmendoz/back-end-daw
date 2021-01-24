package auth

import (
	"crypto/sha256"
	"fmt"
)

func HashPassword(password string) string {
	h := sha256.New()
	_, _ = h.Write([]byte(password))
	hash := fmt.Sprintf("%x", h.Sum(nil))
	return hash
}
