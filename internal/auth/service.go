package auth

import (
	"api-auth/internal/security"
	"api-auth/internal/shared/database"
	"api-auth/internal/user"

	"errors"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func Token(req LoginRequest) (*LoginResponse, error) {
	user := user.User{}

	err := database.DB.
		Debug().
		Preload("Roles").
		Where("username = ?", req.Username).
		First(&user).Error

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

// Apenas valida se o token é meu
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
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("algoritmo inválido")
			}
			return []byte(j.Secret), nil
		},
	)

	return err == nil && token.Valid
}

// Valida e extrai as Claims a partir do token;
func (j *JWTService) ValidateAndExtract(tokenStr string) (*JWTClaims, error) {
	//Se o token estiver expirado, já retorna erro
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&JWTClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("algoritmo inválido")
			}
			return []byte(j.Secret), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, errors.New("token inválido")
	}

	return claims, nil
}
