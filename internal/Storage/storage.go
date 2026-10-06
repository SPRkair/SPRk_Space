package storage

import (
	"SPRk_Space/internal/config"
	"SPRk_Space/internal/logger"
)

type Storage struct {
	Cache *Cache
	DB    *DB
}

func NewStorage(cfg *config.Storage, log logger.Logger) *Storage {
	cache := newCache(&cfg.Cache, log)
	db := newDB(&cfg.DB, log)
	storage := Storage{
		Cache: cache,
		DB:    &db,
	}
	return &storage
}
