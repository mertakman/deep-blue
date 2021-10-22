package main

import (
	"flag"

	"github.com/mertakman/deep-blue/src/api"
)

var host, port string

func init() {
	flag.StringVar(&host, "host", "localhost", "Host name of HTTP Server")
	flag.StringVar(&port, "port", "8080", "Port number of HTTP Server")

	flag.Parse()
}

func main() {
	server := api.NewWebServer(host, port)
	server.Initialize() //  Initializing routers

	if err := server.Run(); err != nil {
		panic(err)
	}
}
