package email

import (
	"github.com/resend/resend-go/v2"

	"zuqui/internal"
)

type Service interface {
	SendEmail(email *EmailRequest) (*EmailSent, error)
}

type service struct {
	client *resend.Client
}

func New(client *resend.Client) *service {
	return &service{client}
}

type EmailRequest struct {
	From        string   `validate:"required"`
	To          []string `validate:"required"`
	Subject     string   `validate:"required"`
	Bcc         []string
	Cc          []string
	ReplyTo     string
	Html        string `validate:"required_without=Text"`
	Text        string `validate:"required_without=Html"`
	Headers     map[string]string
	ScheduledAt string
}

type EmailSent struct {
	Id string
}

func (s *service) SendEmail(email *EmailRequest) (*EmailSent, error) {
	if err := internal.Validate.Struct(email); err != nil {
		return nil, err
	}

	sent, err := s.client.Emails.Send(&resend.SendEmailRequest{
		From:        email.From,
		To:          email.To,
		Subject:     email.Subject,
		Bcc:         email.Bcc,
		Cc:          email.Cc,
		ReplyTo:     email.ReplyTo,
		Html:        email.Html,
		Text:        email.Text,
		Headers:     email.Headers,
		ScheduledAt: email.ScheduledAt,
	})
	if err != nil {
		return nil, err
	}

	return &EmailSent{
		Id: sent.Id,
	}, nil
}
