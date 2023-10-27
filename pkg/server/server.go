package server

import (
	"context"
	"log"
	"net/http"

	"github.com/nikandpro/kafka-fio-server/pkg/config"
	"github.com/nikandpro/kafka-fio-server/pkg/database"
)

type Server struct {
	ctx context.Context
	db  database.Database
	cfg config.Config
}

func NewServer(ctx context.Context, db database.Database, cfg config.Config) *Server {
	return &Server{ctx: ctx, db: db, cfg: cfg}
}

func (s *Server) StartServer() error {
	log.Print("start server")
	http.Handle("/", s.InitRoutes())
	http.ListenAndServe(":8081", nil)

	return nil
}
