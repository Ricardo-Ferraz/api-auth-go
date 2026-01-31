package auth

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var j JWTService

func NewJWTService(secret string) JWTService {

	j.Secret = []byte(secret)
	j.Expiration = 10 * time.Minute

	return j
}

func Generate(id int64, username string, roles []string) (string, int64) {
	now := time.Now()
	expiresAt := now.Add(j.Expiration)

	claims := jwt.MapClaims{
		"sub":      strconv.FormatInt(id, 10),
		"username": username,
		"iat":      now.Unix(),
		"exp":      expiresAt.Unix(),
		"roles":    roles,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := token.SignedString(j.Secret)

	return signedToken, expiresAt.Unix()
}
