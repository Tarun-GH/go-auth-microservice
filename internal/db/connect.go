package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func Connect() *pgx.Conn {
	//constrr := "postgres://postgres:pass123@localhost:5432/go_learning" ---- Main connection string
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:pass123@localhost:5432/go_learning")
	if err != nil {
		log.Fatal("Unable to connect to database", err)
	}

	//Connection check
	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatal("Unable to ping the database", err)
	}
	log.Printf("Connected successfully to database")

	return conn
}
