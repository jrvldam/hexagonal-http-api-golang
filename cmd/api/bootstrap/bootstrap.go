package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jrvldam/hexagonal-http-api-golang/internal/creating"
	"github.com/jrvldam/hexagonal-http-api-golang/internal/platform/bus/inmemory"
	"github.com/jrvldam/hexagonal-http-api-golang/internal/platform/server"
	"github.com/jrvldam/hexagonal-http-api-golang/internal/platform/storage/mysql"
)

const (
	host            = "localhost"
	port            = 8080
	shutdownTimeout = 10 * time.Second

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

	ctx, srv := server.New(context.Background(), host, port, shutdownTimeout, commandBus)
	return srv.Run(ctx)
}
