package auth

import "time"

type jwtService struct {
	secret     []byte
	expiration time.Duration
}
