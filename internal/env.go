package internal

import (
	"fmt"
	"os"
	"strings"

	"github.com/caarlos0/env/v11"
)

type Environment struct {
	PORT    int    `env:"PORT,required"`
	APP_ENV string `env:"APP_ENV,required"`

	OTP_SECRET string `env:"OTP_SECRET,required"`

	GEMINI_API_KEY string `env:"GEMINI_API_KEY,required"`

	RESEND_API_KEY string `env:"RESEND_API_KEY,required"`

	UPSTASH_REDIS_URI string `env:"UPSTASH_REDIS_URI,required"`

	DATABASE_URL string `env:"DATABASE_URL,required"`
}

var Env Environment

func init() {
	red := func(s string) string {
		return "\033[31m" + s + "\033[0m"
	}

	if err := env.Parse(&Env); err != nil {
		fmt.Println(red("Invalid environment variables"))
		fmt.Println(red("  " + strings.ReplaceAll(err.Error()[5:], "; ", "\n  ")))
		os.Exit(1)
	}
}
