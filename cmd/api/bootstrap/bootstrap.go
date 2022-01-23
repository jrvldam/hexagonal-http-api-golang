package bootstrap

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jrvldam/hexagonal-http-api-golang/internal/platform/server"
	"github.com/jrvldam/hexagonal-http-api-golang/internal/platform/storage/mysql"
)

const (
	host = "localhost"
	port = 8080

	dbUser = "dbUser"
	dbPass = "dbPass"
	dbHost = "dbHost"
	dbPort = "dbPort"
	dbName = "dbName"
)

func Run() error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)

	if err != nil {
		return err
	}

	courseRepository := mysql.NewCourseRepository(db)

	srv := server.New(host, port, courseRepository)
	return srv.Run()
}
