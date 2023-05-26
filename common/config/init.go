package config

import "github.com/caarlos0/env/v8"

type (
	config struct {
		Port            int    `env:"PORT,unset" envDefault:"4001"`
		DatabaseURL     string `env:"DATABASE_URL,unset"`
		JWTAccessToken  string `env:"ACCESS_TOKEN,unset"`
		XenditSecretKey string `env:"XENDIT_SECRET_KEY,unset"`
	}
)

func LoadConfig() *config {
	cfg := &config{}
	opts := env.Options{RequiredIfNoDef: true}
	if err := env.ParseWithOptions(cfg, opts); err != nil {
		panic(err)
	}
	return cfg
}
