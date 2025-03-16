package auth

import (
	"context"
	"errors"
	"math/rand/v2"
	"time"
)

func (s *service) CreateOTP(key string) (string, error) {
	if key == "" {
		return "", errors.New("otp key is required")
	}
	otp := createOTP()

	return otp, s.client.Set(context.Background(), key, otp, 10*time.Minute).Err()
}

func (s *service) VerifyOTP(key string, otp string) bool {
	entry := s.client.Get(context.Background(), key)
	if entry.Err() != nil {
		return false
	}

	verified := otp == entry.Val()
	if verified {
		s.client.Del(context.Background(), key)
	}

	return verified
}

func createOTP() string {
	chars := "ABCDEFGHJKLMNPQRTUVWXY346789" // Without O, 0, I, 1, S, 5, Z, 2
	r := make([]byte, 6)
	for i := range r {
		r[i] = chars[rand.IntN(len(chars))]
	}

	return string(r)
}
