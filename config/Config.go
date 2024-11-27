package config

import (
	"github.com/gin-gonic/gin"
	"github.com/jmscatena/Fatec_Sert_SGLab/routes"
	"log"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer(port string) Server {
	return Server{
		port:   port,
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)
	log.Printf("Server running at port: %v", s.port)
	log.Fatal(router.Run(":" + s.port))
}
