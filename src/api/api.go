package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type HTTPServer struct {
	app        *fiber.App
	host, port string
}

func NewWebServer(host, port string) *HTTPServer {
	server := new(HTTPServer)
	server.app = fiber.New()
	server.port = port
	server.host = host

	server.app.Use(cors.New())
	server.app.Use(compress.New())
	server.app.Use(recover.New())
	server.app.Use(cors.New(cors.Config{
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,OPTION",
		AllowHeaders: "Origin,Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token",
	}))

	return server
}

func (server *HTTPServer) Initialize() {
	server.app.Get("/deep-blue/knight/calculate-route", CalculateRoute)
	// #todo : write http unit tests.
}

func (server *HTTPServer) Run() error {
	return server.app.Listen(fmt.Sprintf("%s:%s", server.host, server.port))
}
