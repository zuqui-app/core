package email

import (
	"bytes"
	"html/template"

	"github.com/resend/resend-go/v2"

	"zuqui-core/internal/services"
)

type SendOTPProps struct {
	Username string
	OTP      string
}

func SendOTPEmail(email string, props SendOTPProps) (sent *resend.SendEmailResponse, error error) {
	temp, err := template.ParseFiles("internal/email/otp_template.html")
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err := temp.Execute(&buf, props); err != nil {
		return nil, err
	}

	html := buf.String()
	params := &resend.SendEmailRequest{
		From:    "Zuqui <zuqui@nichtsam.com>",
		To:      []string{email},
		Subject: "🔑 Here's your OTP for Zuqui!",
		Html:    html,
	}

	sent, err = services.ResendClient.Emails.Send(params)
	if err != nil {
		return nil, err
	}

	return sent, nil
}
