package main

import (
	"context"
	"log"
	"net/http"

	"github.com/Tarun-GH/go-rest-microservice/internal/db"
	internal "github.com/Tarun-GH/go-rest-microservice/internal/handlers"
)

func main() {
	conn := db.Connect()
	defer conn.Close(context.Background()) //context is used for the context of duration or conditions for it to last/ trigger

	var name string
	err := conn.QueryRow(context.Background(), "SELECT name FROM test_items WHERE id=2").Scan(&name)
	if err != nil {
		log.Println("Query failed:", err)
	} else {
		log.Println("Name from DB:", name)
	}

	// Inserting a user
	err = repository.InsertUser(conn, "John Doe", "johndoe@example.com", "hashed_password123")
	if err != nil {
		log.Fatal("Error inserting user:", err)
	}

	http.HandleFunc("/health", internal.Health)
	http.HandleFunc("/version", internal.Version)

	log.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Couldn't start the server", err)
	}
}
