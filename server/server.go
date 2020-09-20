package server

import (
	"context"
	"net/http"

	"github.com/geraldofigueiredo/mypage/config"
	"github.com/geraldofigueiredo/mypage/service"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start() error {
	// load api configs
	config := config.LoadConfig()
	// services
	services := service.NewServices(config)

	ctx := context.Background()
	ctx = context.WithValue(ctx, "config", config)
	ctx = context.WithValue(ctx, "services", services)

	services.Log.Infof("server running on port %s", config.AppPort)
	return http.ListenAndServe(config.AppPort, nil)
}
