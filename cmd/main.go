package main

import "pasour/internal/interfaces"

func main() {
	server := interfaces.Bootstrap()
	server.SetupRoutes()
	server.Start(":5000")
}
