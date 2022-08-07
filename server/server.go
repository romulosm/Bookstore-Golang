package server

import (
	"livraria-golang/server/routes"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{port: "3000", server: gin.Default()}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)
	log.Print("Server started in port: ", s.port)
	log.Fatal(router.Run(":" + s.port))
}
