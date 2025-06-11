package manager

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"log"
	"time"
)

var (
	ErrSigningMethod      = errors.New("invalid signing method")
	ErrTokenExpired       = errors.New("token is expired")
	ErrInvalidClaimFormat = errors.New("invalid claim format")
	ErrSubClaimNotFound   = errors.New("sub claim not found")
	ErrInvalidSubFormat   = errors.New("invalid sub format")
)

type Manager interface {
	ParseToken(accessToken string) (uint64, error)
	NewAccessToken(userID uint64, ttl time.Duration) (string, error)
	NewRefreshToken() uuid.UUID
}

type TokenManager struct {
	signingKey string
}

func MustLoadTokenManager(secret string) Manager {
	if secret == "" {
		log.Fatalln("signingKey is empty")
	}
	return &TokenManager{signingKey: secret}
}

func (t TokenManager) ParseToken(jwtToken string) (uint64, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrSigningMethod
		}

		return []byte(t.signingKey), nil
	})

	if err != nil {
		return 0, ErrTokenExpired
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, ErrInvalidClaimFormat
	}

	sub, ok := claims["sub"]
	if !ok {
		return 0, ErrSubClaimNotFound
	}

	subFloat, ok := sub.(float64)
	if !ok {
		return 0, ErrInvalidSubFormat
	}

	return uint64(subFloat), nil
}

func (t TokenManager) NewAccessToken(userID uint64, ttl time.Duration) (string, error) {
	claim := jwt.MapClaims{
		"exp": time.Now().UTC().Add(ttl).Unix(),
		"sub": userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	return token.SignedString([]byte(t.signingKey))
}

func (t TokenManager) NewRefreshToken() uuid.UUID {
	return uuid.New()
}
