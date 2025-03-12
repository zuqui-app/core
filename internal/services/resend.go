package services

import (
	"github.com/resend/resend-go/v2"

	"zuqui-core/internal"
)

var ResendClient *resend.Client

func init() {
	ResendClient = resend.NewClient(internal.Env.RESEND_API_KEY)
}
