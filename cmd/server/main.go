package main 

import (
	"log"
	"github.com/GGiecold/go_distributed_systems/internal/server"
)

func main() {
	server := server.NewHttpServer(":8080")
	log.Fatal(server.ListenAndServe())
}

