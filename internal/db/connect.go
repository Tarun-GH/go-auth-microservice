package db

import (
	"context"
	"log"

	"github.com/Tarun-GH/go-rest-microservice/internal/config"
	"github.com/jackc/pgx/v5"
)

func Connect(cfg *config.Config) *pgx.Conn { //cfg env dependency hai, should be handled that way
	constrr := ("postgres://" + cfg.DBUser + ":" + cfg.DBPassword + "@" + cfg.DBHost + ":" + cfg.DBPort + "/" + cfg.DBName)

	conn, err := pgx.Connect(context.Background(), constrr)
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
