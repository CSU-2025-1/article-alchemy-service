package hasher

import (
	"crypto/sha256"
	"encoding/hex"
	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

func HashPassword(password string, salt []byte) string {
	return hex.EncodeToString(argon2.IDKey([]byte(password), salt, 1, 32*1024, 2, 32))
}

func NewBcryptHash(text string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.MinCost)
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
