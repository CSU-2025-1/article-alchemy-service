package hasher

import (
	"crypto/sha256"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

func NewBcryptHash(text string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(passwordHash), nil
}

func CompareBcryptHash(hash string, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
	if err != nil {
		return false
	}

	return true
}

func GenerateHash(items ...string) string {
	var data strings.Builder
	for _, item := range items {
		data.WriteString(item)
	}

	hash := sha256.Sum256([]byte(data.String()))

	return hex.EncodeToString(hash[:])
}
