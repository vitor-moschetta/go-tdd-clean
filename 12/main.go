package main

import (
	"go-tdd-clean/12/infrastructure/api/server"
	"log"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
