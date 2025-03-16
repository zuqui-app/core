package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

type Service interface {
	CreateOTP(key string) (string, error)
	VerifyOTP(key string, otp string) bool
	CreateAccessToken(subject string) (string, error)
	CreateRefreshToken(subject string) (string, error)
	CreateTokenPair(subject string) (access string, refresh string, error error)
	VerifyAccessToken(tokenString string) (*jwt.RegisteredClaims, error)
	VerifyRefreshToken(tokenString string) (*jwt.RegisteredClaims, error)
	RevokeRefreshToken(tokenId string) error
}

type service struct {
	client *redis.Client
}

func New(client *redis.Client) Service {
	return &service{client}
}
