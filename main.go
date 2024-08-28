package main

import (
	"log"
)

func main() {
	server, err := InitializeServer()
	if err != nil {
		log.Fatalf("failed to initialize server: %v", err)
	}
	server.Run()
}
