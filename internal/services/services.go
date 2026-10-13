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
	log     logger.Logger
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
		log:     log,
		Storage: storage,
		FakeService: &FakeService{
			ctx:   ctx,
			log:   log,
			Users: storage.User,
		},
	}
	return &services
}

func (s Services) Run() {
	s.log.Println("Info", "Services Run...")
	s.FakeService.Run()
	s.log.Println("Info", "Services Run Complete.")
}
