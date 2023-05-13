package main

import (
	"log"

	"github.com/ryanralphs/distributed-service/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
