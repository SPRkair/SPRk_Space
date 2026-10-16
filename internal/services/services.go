package services

import (
	"SPRk_Space/internal/config"
	"SPRk_Space/internal/logger"
	"SPRk_Space/internal/storage"
	"context"
)

type Services struct {
	ctx     context.Context
	Config  *config.Services
	logger  logger.Logger
	Storage *storage.Storage

	CacheSync   *CacheSync
	Executor    *Executor
	FakeService *FakeService
}

type CacheSync struct {
}

type Executor struct {
}

func NewServices(ctx context.Context, cfg *config.Services, log logger.Logger, storage *storage.Storage) *Services {
	services := Services{
		Config:  cfg,
		logger:  log,
		Storage: storage,
		FakeService: &FakeService{
			ctx:    ctx,
			logger: log,
			Users:  storage.User,
		},
	}
	return &services
}

func (s Services) Run() {
	s.logger.Info("Services Run...")
	s.FakeService.Run()
	s.logger.Info("Services Run Complete.")
}

func (s Services) Close() {
	s.logger.Debug_info("Services Close")
}
