package main

import (
	"log"

	"github.com/jrvldam/hexagonal-http-api-golang/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
