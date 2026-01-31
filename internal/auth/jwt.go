package auth

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var j jwtService

func NewJWTService(secret string) {

	j.secret = []byte(secret)
	j.expiration = 10 * time.Minute
}

func Generate(id int64, username string, roles []string) (string, int64) {
	now := time.Now()
	expiresAt := now.Add(j.expiration)

	claims := jwt.MapClaims{
		"sub":      strconv.FormatInt(id, 10),
		"username": username,
		"iat":      now.Unix(),
		"exp":      expiresAt.Unix(),
		"roles":    roles,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := token.SignedString(j.secret)

	return signedToken, expiresAt.Unix()
}
