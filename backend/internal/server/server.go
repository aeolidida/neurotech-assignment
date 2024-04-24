package server

import (
	"neurotech-assignment/backend/internal/config"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	host   string
	port   int
}

func NewServer(cfg config.HTTPServer, repo PatientRepository) *Server {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "OPTIONS"}
	router.Use(cors.New(config))

	api := NewPatientAPI(repo)
	api.RegisterHandlers(router)

	return &Server{
		router: router,
		host:   cfg.Host,
		port:   cfg.Port,
	}
}

func (s *Server) Serve() error {
	addr := s.host + ":" + strconv.Itoa(s.port)
	err := s.router.Run(addr)
	if err != nil {
		return err
	}
	return nil
}
