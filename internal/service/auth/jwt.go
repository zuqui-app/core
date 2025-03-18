package auth

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/oklog/ulid/v2"

	"zuqui/internal"
)

const (
	access_token_ttl  = 1 * time.Hour
	refresh_token_ttl = 30 * 24 * time.Hour
)

var (
	method     = jwt.SigningMethodHS256
	signingKey = []byte(internal.Env.JWT_SECRET)
)

func (s *service) CreateAccessToken(subject string) (string, error) {
	claims := jwt.RegisteredClaims{
		ID:        ulid.Make().String(),
		Issuer:    "Zuqui",
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   subject,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(access_token_ttl)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signingKey)
}

func (s *service) CreateRefreshToken(subject string) (string, error) {
	claims := jwt.RegisteredClaims{
		ID:        ulid.Make().String(),
		Issuer:    "Zuqui",
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   subject,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(refresh_token_ttl)),
	}
	token := jwt.NewWithClaims(method, claims)
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}

	entry := s.client.Set(context.Background(), "refresh_token:"+claims.ID, nil, refresh_token_ttl)
	err = entry.Err()
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *service) CreateTokenPair(subject string) (string, string, error) {
	access, err := s.CreateAccessToken(subject)
	if err != nil {
		return "", "", err
	}
	refresh, err := s.CreateRefreshToken(subject)
	if err != nil {
		return "", "", err
	}

	return access, refresh, nil
}

func (s *service) VerifyAccessToken(tokenString string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&jwt.RegisteredClaims{},
		func(t *jwt.Token) (any, error) {
			if t.Method.Alg() != method.Alg() {
				return nil, NewAuthError("Unexpected Algorithm")
			}
			return signingKey, nil
		},
	)
	if err != nil {
		return nil, NewAuthError("Invalid Token").withError(err)
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return nil, NewAuthError("Unexpected Claims")
	}

	return claims, nil
}

func (s *service) VerifyRefreshToken(tokenString string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&jwt.RegisteredClaims{},
		func(t *jwt.Token) (any, error) {
			if t.Method.Alg() != method.Alg() {
				return nil, NewAuthError("Unexpected Algorithm")
			}
			return signingKey, nil
		},
	)
	if err != nil {
		return nil, NewAuthError("Invalid Token").withError(err)
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return nil, NewAuthError("Unexpected Claims")
	}

	entry := s.client.Exists(context.Background(), "refresh_token:"+claims.ID)
	err = entry.Err()
	if err != nil {
		return nil, err
	}

	if entry.Val() == 0 {
		return nil, NewAuthError("Invalid Refresh Token")
	}

	return claims, nil
}

func (s *service) RevokeRefreshToken(tokenId string) error {
	entry := s.client.Del(context.Background(), "refresh_token:"+tokenId)
	return entry.Err()
}
