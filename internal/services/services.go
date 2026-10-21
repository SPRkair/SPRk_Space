package services

import (
	"SPRk_Space/internal/config"
	"SPRk_Space/internal/logger"
	"SPRk_Space/internal/services/auth"
	"SPRk_Space/internal/storage"
	"context"
)

type Services struct {
	ctx    context.Context
	config *config.Services
	logger logger.Logger

	FakeService *FakeService
	Auth        *auth.Auth
}

func NewServices(ctx context.Context, cfg *config.Services, log logger.Logger, storage *storage.Storage) *Services {
	services := Services{
		config:      cfg,
		logger:      log,
		FakeService: NewFakeService(ctx, &cfg.FakeService, log, storage),
		Auth:        auth.NewAuth(ctx, &cfg.Auth, log, storage),
	}
	return &services
}

func (s *Services) Run() {
	s.logger.Debug_info("Services Run...")
	s.FakeService.Run()
	s.logger.Debug_info("Services Run Complete.")
}

func (s *Services) Close() {
	s.logger.Debug_info("Services Close")
	s.Auth.Close()
	s.FakeService.Close()
}
