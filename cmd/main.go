package main

import (
	"github.com/lucasmbrute2/reverse-proxy-from-scratch/internal/server"
	"log"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatalf("could not start the server: %v", err)
	}
}
