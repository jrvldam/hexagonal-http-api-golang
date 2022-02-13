package bootstrap

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jrvldam/hexagonal-http-api-golang/internal/creating"
	"github.com/jrvldam/hexagonal-http-api-golang/internal/platform/bus/inmemory"
	"github.com/jrvldam/hexagonal-http-api-golang/internal/platform/server"
	"github.com/jrvldam/hexagonal-http-api-golang/internal/platform/storage/mysql"
)

const (
	host = "localhost"
	port = 8080

	dbUser = "root"
	dbPass = "BATMAN"
	dbHost = "127.0.0.1"
	dbPort = "8083"
	dbName = "courses_db"
)

func Run() error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)

	if err != nil {
		return err
	}

	courseRepository := mysql.NewCourseRepository(db)
	creatingCourseService := creating.NewCourseService(courseRepository)
	createCourseCommandHandler := creating.NewCourseCommandHandler(creatingCourseService)

	var commandBus = inmemory.NewCommandBus()

	commandBus.Register(creating.CourseCommandType, createCourseCommandHandler)

	srv := server.New(host, port, commandBus)
	return srv.Run()
}
