package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jrvldam/hexagonal-http-api-golang/internal/platform/server/handler/courses"
	"github.com/jrvldam/hexagonal-http-api-golang/internal/platform/server/handler/health"
	"github.com/jrvldam/hexagonal-http-api-golang/kit/command"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	// deps
	commandBus command.Bus
}

func New(host string, port uint, commandBus command.Bus) Server {
	srv := Server{
		httpAddr: fmt.Sprintf("%s:%d", host, port),
		engine:   gin.New(),

		commandBus: commandBus,
	}

	srv.registerRoutes()

	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)

	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
	s.engine.POST("/courses", courses.CreateHandler(s.commandBus))
}
