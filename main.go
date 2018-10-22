package main

import "./server"
import "./server/ServerConfig"

func main() {
	cfg := ServerConfig.New(8080, "0.0.0.0")
	s := server.New(*cfg)
	s.Connect()
}
