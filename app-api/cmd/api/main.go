package main

import (
	"flag"
	"log"
	"task-api/internal/database"
	"task-api/internal/server"
)

var (
	listenAddr   = flag.String("address", "localhost:8082", "address to listen")
	databaseAddr = flag.String("database-url", "db", "database address")
	databasePort = flag.String("database-port", "5432", "database port")
)

func main() {

	flag.Parse()

	pg := database.NewPostgres()
	serv := server.NewServer(pg)

	serv.InitSever()
	if err := pg.Connect(*databaseAddr, *databasePort); err != nil {
		log.Fatalf("database error: %v", err)
	}

	log.Fatal(serv.Listen(*listenAddr))
}
