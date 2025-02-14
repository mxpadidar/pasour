package interfaces

import (
	"log"
	"pasour/internal/application"
	"pasour/internal/application/services"
	"pasour/internal/infrastracture"
	"pasour/internal/infrastracture/sqlrepo"
)

func Bootstrap() application.Server {
	db, err := infrastracture.NewDB()

	if err != nil {
		log.Fatal(err)
	}
	userRepo := sqlrepo.NewSqlUserRepo(db)
	userService := services.NewUserService(userRepo)
	server := NewHttpServer(userService)
	return server
}
