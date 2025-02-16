package interfaces

import (
	"log"
	"net/http"
	"pasour/internal/domain/services"
	"pasour/internal/interfaces/handlers"
)

type HttpServer struct {
	userSrv  services.UserService
	TokenSrv services.TokenService
	Router   *http.ServeMux
}

func NewHttpServer(userSrv services.UserService, tokenSrv services.TokenService) *HttpServer {
	return &HttpServer{
		userSrv:  userSrv,
		TokenSrv: tokenSrv,
		Router:   http.NewServeMux(),
	}
}

func (s *HttpServer) SetupRoutes() {
	userHandler := handlers.NewUserHandler(s.userSrv, s.Router)
	userHandler.SetupRoutes()
}

func (s *HttpServer) Start(addr string) {
	log.Printf("Starting server on %s", addr)
	if err := http.ListenAndServe(addr, s.Router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
