package auth

import (
	"fmt"
	"time"
)

type AuthError struct {
	message string
	err     error
}

func NewAuthError(msg string) *AuthError {
	return &AuthError{
		message: msg,
	}
}

func (e *AuthError) withError(err error) *AuthError {
	e.err = err
	return e
}

func (e *AuthError) Error() string {
	if e.err != nil {
		return e.message + "\n" + e.err.Error()
	}
	return e.message
}

type CooldownError struct {
	Seconds int
}

func NewCooldownError(t time.Duration) *CooldownError {
	return &CooldownError{int(t.Seconds())}
}

func (e *CooldownError) Error() string {
	return fmt.Sprintf("OTP is on cooldown. You can try again in %ds.", e.Seconds)
}
