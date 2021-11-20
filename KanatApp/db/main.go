package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "db"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "junction21"
)
func dbConnect(dbInfo string) *sql.DB{
	db, err := sql.Open("postgres", dbInfo)
	for err != nil {
		sql.Open("postgres", dbInfo)
	}
	return  db
}
func dbQueryHandler(db *sql.DB, query string){
	_, err := db.Exec(query)
	if err != nil{
		fmt.Println(query)
		panic(err)
	}
}
func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s sslmode=disable",
		host, port, user, password)
	db := dbConnect(psqlInfo)
	query := "CREATE DATABASE "+ dbname
	dbQueryHandler(db, query)
	db.Close()


	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db = dbConnect(psqlInfo)


	query  = "CREATE TABLE users ( " +
		"id integer GENERATED ALWAYS AS IDENTITY, " +
		"name varchar(32) UNIQUE, " +
		"rating integer, " +
		"photo bytea, " +
		"PRIMARY KEY(id))"
	dbQueryHandler(db, query)

	query = "CREATE TABLE tasks ( " +
		"id integer GENERATED ALWAYS AS IDENTITY, " +
		"user_id integer, " +
		"name varchar(32), " +
		"description varchar(32), " +
		"start_date date, " +
		"end_date date, " +
		"PRIMARY KEY(id), " +
		"CONSTRAINT fk_users FOREIGN KEY(user_id) REFERENCES users(id))"

	dbQueryHandler(db, query)
	db.Close()
	fmt.Println("Successfully connected!")
}