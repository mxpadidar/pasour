package interfaces

import (
	"log"

	"pasour/internal/application"
	"pasour/internal/application/services"
	"pasour/internal/infrastracture/configs"
	"pasour/internal/infrastracture/db"
	"pasour/internal/infrastracture/sqlrepo"
)

func Bootstrap() application.Server {
	db, err := db.NewDB()

	if err != nil {
		log.Fatal(err)
	}
	userRepo := sqlrepo.NewSqlUserRepo(db)
	userService := services.NewUserService(userRepo)
	tokenSrv := services.NewTokenService(configs.Configs.Secret, configs.Configs.JwtDuration)
	server := NewHttpServer(userService, tokenSrv)
	return server
}
