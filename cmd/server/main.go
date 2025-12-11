package main

import (
	"context"
	"log"
	"net/http"

	"github.com/Tarun-GH/go-rest-microservice/internal/db"
	"github.com/Tarun-GH/go-rest-microservice/internal/handlers"
	"github.com/Tarun-GH/go-rest-microservice/internal/routes"
	"github.com/go-chi/chi/v5"
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
	handlers.DB = conn

	/*	// Inserting a user
		err = repository.InsertUser(conn, "hoho", "hoho@example.com", "h_psd13", "users")
		if err != nil {
			log.Fatal("Error inserting user:", err)
		}

		// Get User by email
		user, err := repository.GetUserByEmail(conn, "hoho@example.com", "users")
		if err != nil {
			log.Fatal("Couldn't get User", err)
		}
		log.Printf("User details\n :%s\n :%s\n :%s\n", user.Name, user.Email, user.CreatedAt)
	*/

	//Handling routes/endpoint
	r := chi.NewRouter()
	routes.RegisterRoutes(r)

	r.Get("/health", handlers.Health) //----replaced http.HandleFunc("",)
	r.Get("/version", handlers.Version)

	log.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Couldn't start the server", err)
	}
}
