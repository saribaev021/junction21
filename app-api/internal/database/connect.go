package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgresDB struct {
	db *sql.DB
}

func NewPostgres() *PostgresDB {
	return &PostgresDB{}
}

func (pg *PostgresDB) Connect(host, port string) error {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable dbname=junction21",
		host, port, "postgres", "password"))
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	pg.db = db

	return nil
}
