package server

import (
	"./ServerConfig"
	"fmt"
)

type Server struct {
	serverConfig ServerConfig.ServerConfig
}

func New(config ServerConfig.ServerConfig) *Server {
	return &Server{config}
}

func (server *Server) Connect() bool {
	fmt.Printf("Connecting to Port: %d Host: %s", server.serverConfig.Port, server.serverConfig.Host)
	return true
}

func (server *Server) Disconnect() bool {
	fmt.Println("Disconnecting server")
	return true
}
