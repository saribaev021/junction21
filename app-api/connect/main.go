package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "db"
	port     = 5432
	user     = "postgres"
	password = "password"
)
func dbConnect(dbInfo string){
	db, err := sql.Open("postgres", dbInfo)
	if err != nil{
		fmt.Println("err:", err)
	}

	for err = db.Ping(); err != nil;{
		fmt.Println("wait...", err)
		time.Sleep(time.Second)
	}
	db.Close()
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s sslmode=disable",
		host, port, user, password)
	dbConnect(psqlInfo)
}