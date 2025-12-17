package config

import "os"

type Config struct {
	JWTSecret       string
	AccessTokenTTL  int64
	RefreshTokenTTL int64
	Issuer          string
}

func LoadConfig() *Config {
	return &Config{
		JWTSecret:       os.Getenv("JWT_SECRET"),
		AccessTokenTTL:  15 * 60,
		RefreshTokenTTL: 7 * 24 * 3600,
		Issuer:          "shop-api",
	}
}
