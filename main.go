package main

import (
	"api-auth/internal/auth"
	"api-auth/internal/config"
	"api-auth/internal/router"
	"api-auth/internal/shared/database"
)

func main() {

	jwtCfg := config.LoadJWT()

	j := auth.NewJWTService(jwtCfg.Secret)

	database.ConectaComBancoDeDados()

	router.SetupRouter(j).Run()
	// Inicia o servidor na porta 8080
	// escuta e serve em 0.0.0.0:8080
}
