package services

import (
	"SPRk_Space/internal/config"
	"SPRk_Space/internal/logger"
	"SPRk_Space/internal/storage"
)

type Services struct {
	Config  *config.Services
	log     logger.Logger
	Storage *storage.Storage

	CacheSync *CacheSync
	Executor  *Executor
}

type CacheSync struct {
}

type Executor struct {
}

func NewServices(cfg *config.Services, log logger.Logger, storage *storage.Storage) *Services {
	services := Services{
		Config:  cfg,
		log:     log,
		Storage: storage,
	}
	return &services
}

func (s Services) Run() {

}
