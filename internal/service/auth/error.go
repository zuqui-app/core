package auth

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
