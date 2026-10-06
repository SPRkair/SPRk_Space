package storage

import (
	"SPRk_Space/internal/config"
	"SPRk_Space/internal/logger"
)

type Storage struct {
	cache *Cache
	db    *DB

	FakeStorage FakeStorage
}

func NewStorage(cfg *config.Storage, log logger.Logger) *Storage {
	cache := newCache(&cfg.Cache, log)
	db := newDB(&cfg.DB, log)
	storage := Storage{
		cache: cache,
		db:    &db,
	}
	return &storage
}

func (s Storage) Get() interface{} {

	return nil
}
