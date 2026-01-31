package config

import (
	"os"
	"strconv"
)

type JWTConfig struct {
	Secret string
	TTL    int64 // minutos
}

func LoadJWT() JWTConfig {
	ttl, _ := strconv.ParseInt(
		"0",
		10,
		64,
	)

	if ttl == 0 {
		ttl = 60
	}

	return JWTConfig{
		Secret: os.Getenv("AAAAAAAAAA_SECRET"),
		TTL:    ttl,
	}
}
