package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	httpAddr = ":8080"
)

func main() {
	fmt.Println("Serve running at:", httpAddr)

	srv := gin.New()
	srv.GET("/health", healthHandler)

	log.Fatal(srv.Run(httpAddr))
}

func healthHandler(c *gin.Context) {
	c.String(http.StatusOK, "everything is 0k")
}
