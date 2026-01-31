package auth

import (
	"api-auth/internal/security"
	"api-auth/internal/shared/database"
	"api-auth/internal/user"
	"errors"
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func Token(req LoginRequest) (*LoginResponse, error) {
	var user user.User

	err := database.DB.
		Debug().
		Preload("Roles").
		Where("username = ?", req.Username).
		First(&user).Error

	fmt.Println(user)

	if err != nil {
		return nil, ErrInvalidCredentials
	}

	if !security.CheckPasswordHash(req.Password, user.PassHash) {
		return nil, ErrInvalidCredentials
	}

	token, exp := Generate(
		user.Id,
		user.Username,
		user.ReturnNameRoles(),
	)

	return &LoginResponse{
		AccessToken: token,
		ExpiresIn:   exp,
	}, nil
}

func isMyToken(bearer string) bool {
	if bearer == "" {
		return false
	}

	parts := strings.Split(bearer, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return false
	}

	tokenStr := parts[1]

	token, err := jwt.Parse(
		tokenStr,
		func(token *jwt.Token) (interface{}, error) {
			// Impede troca de algoritmo
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("algoritmo inválido")
			}
			return []byte(j.secret), nil
		},
	)

	return err == nil && token.Valid
}
