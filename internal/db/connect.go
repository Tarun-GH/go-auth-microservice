package db

import (
	"context"
	"log"

	"github.com/Tarun-GH/go-rest-microservice/internal/config"
	"github.com/jackc/pgx/v5"
)

func Connect() *pgx.Conn {
	//constrr := "postgres://postgres:pass123@localhost:5432/go_learning" ---- Main connection string
	cfg := config.Load()
	constrr := ("postgres://" + cfg.DBUser + ":" + cfg.DBPassword + "@" + cfg.DBHost + ":" + cfg.DBPort + "/" + cfg.DBName)

	// log.Println("ENV DB_HOST =", os.Getenv("DB_HOST"))
	// log.Println("CFG DBHost =", cfg.DBHost) //config data check

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
