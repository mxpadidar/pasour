package interfaces

import (
	"log"
	"net/http"
	"pasour/internal/domain/services"
	"pasour/internal/interfaces/handlers"
)

type HttpServer struct {
	UserService services.UserService
	Router      *http.ServeMux
}

func NewHttpServer(userService services.UserService) *HttpServer {
	return &HttpServer{
		UserService: userService,
		Router:      http.NewServeMux(),
	}
}

func (s *HttpServer) SetupRoutes() {
	userHandler := handlers.NewUserHandler(s.UserService, s.Router)
	userHandler.SetupRoutes()
}

func (s *HttpServer) Start(addr string) {
	log.Printf("Starting server on %s", addr)
	if err := http.ListenAndServe(addr, s.Router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
