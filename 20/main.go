package main

import (
	"go-tdd-clean/20/api"
	"log"
)

func main() {
	if err := api.Run(); err != nil {
		log.Fatal(err)
	}
}
