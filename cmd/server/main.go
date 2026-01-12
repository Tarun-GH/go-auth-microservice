package main

import (
	"context"
	"log"
	"net/http"

	"github.com/Tarun-GH/go-rest-microservice/internal/config"
	"github.com/Tarun-GH/go-rest-microservice/internal/db"
	"github.com/Tarun-GH/go-rest-microservice/internal/handlers"
	"github.com/Tarun-GH/go-rest-microservice/internal/queue"
	"github.com/Tarun-GH/go-rest-microservice/internal/routes"
	"github.com/Tarun-GH/go-rest-microservice/internal/utils"
	"github.com/go-chi/chi/v5"
)

func main() {
	cfg := config.Load()
	conn := db.Connect(cfg)
	defer conn.Close(context.Background()) //context is used for the context of duration or 'conditions for it to last/ trigger'

	//Redis initialisation
	redisClient := config.NewRedisClient(cfg.RedisHost)
	utils.InitRedis(redisClient)

	//RabbitMQ connection
	mq, err := queue.NewRabbitMQ(cfg.RabbitMQURL)
	if err != nil {
		log.Fatal("Error connecting to RabbitMQ:", err)
	}
	defer mq.Close()

	//Handling routes/endpoint
	h := handlers.NewHandler(conn, mq, []byte(cfg.JWTSecret)) //both connections made n send to Handler type
	r := chi.NewRouter()
	routes.RegisterRoutes(r, h)

	r.Get("/health", handlers.Health) //----replaced http.HandleFunc("",)
	r.Get("/version", handlers.Version)

	log.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Couldn't start the server", err)
	}
}
