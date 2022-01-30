package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jrvldam/hexagonal-http-api-golang/internal/creating"
	"github.com/jrvldam/hexagonal-http-api-golang/internal/platform/server/handler/courses"
	"github.com/jrvldam/hexagonal-http-api-golang/internal/platform/server/handler/health"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	// deps
	creatingCourseService creating.CourseService
}

func New(host string, port uint, creatingCourseService creating.CourseService) Server {
	srv := Server{
		httpAddr: fmt.Sprintf("%s:%d", host, port),
		engine:   gin.New(),

		creatingCourseService: creatingCourseService,
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
	s.engine.POST("/courses", courses.CreateHandler(s.creatingCourseService))
}
