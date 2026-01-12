package handlers

import (
	"github.com/Tarun-GH/go-rest-microservice/internal/queue"
	"github.com/jackc/pgx/v5"
)

type Handler struct {
	DB        *pgx.Conn
	MQ        *queue.RabbitMQ
	JWTSecret []byte
}

func NewHandler(db *pgx.Conn, mq *queue.RabbitMQ, JWTSecret []byte) *Handler {
	return &Handler{
		DB:        db,
		MQ:        mq,
		JWTSecret: JWTSecret,
	}
}
