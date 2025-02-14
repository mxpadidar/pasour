package application

type Server interface {
	SetupRoutes()
	Start(addr string)
}
