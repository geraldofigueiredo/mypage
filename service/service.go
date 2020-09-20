package service

import (
	"github.com/geraldofigueiredo/mypage/config"
	"github.com/op/go-logging"
)

// Services Struct with all services in the system
type Services struct {
	Log logging.Logger
}

// NewServices Create a new Services
func NewServices(config *config.Config) *Services {
	log := newLogger(config)

	return &Services{
		Log: *log,
	}
}
