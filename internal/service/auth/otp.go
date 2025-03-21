package auth

import (
	"context"
	"errors"
	"math/rand/v2"
	"time"
)

func getOtpKey(key string) string {
	return "otp:" + key
}

func getCooldownKey(key string) string {
	return "otp_cooldown:" + key
}

func (s *service) CreateOTP(key string) (string, error) {
	if key == "" {
		return "", errors.New("otp key is required")
	}

	cooldownKey := getCooldownKey(key)
	entry := s.client.TTL(context.Background(), cooldownKey)
	err := entry.Err()
	if err != nil {
		return "", err
	}

	if ttl := entry.Val(); ttl > 0 {
		return "", NewCooldownError(ttl)
	}

	otp := createOTP()
	otpKey := getOtpKey(key)
	if err := s.client.Set(context.Background(), otpKey, otp, 10*time.Minute).Err(); err != nil {
		return "", err
	}

	if err := s.client.Set(context.Background(), cooldownKey, nil, 1*time.Minute).Err(); err != nil {
		return "", err
	}

	return otp, nil
}

func (s *service) VerifyOTP(key string, otp string) bool {
	otpKey := getOtpKey(key)
	entry := s.client.Get(context.Background(), otpKey)
	if entry.Err() != nil {
		return false
	}

	verified := otp == entry.Val()
	if verified {
		s.client.Del(context.Background(), otpKey)
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
