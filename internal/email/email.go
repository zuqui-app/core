package email

import "github.com/resend/resend-go/v2"

type EmailService struct {
	client *resend.Client
}

func New(client *resend.Client) *EmailService {
	return &EmailService{client}
}
