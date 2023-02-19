package main

import (
	"log"
	"marketplace/api"

	"os"
)

func main() {
	// Get the server address from the command-line arguments
	args := os.Args
	if len(args) != 2 {
		log.Fatal("Usage: go run cmd/backend/main.go <server-address>")
	}
	serverAddr := args[1]

	api.StartServer(serverAddr)
}
