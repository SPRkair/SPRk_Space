package storage

import (
	"SPRk_Space/internal/config"
	"SPRk_Space/internal/entities"
	"SPRk_Space/internal/logger"
)

type Cache struct {
	config *config.Cache
	logger *logger.Logger

	UserCache *UserCache
}

type UserCache struct {
	Users map[int64]*entities.User
}

func newCache(cfg *config.Cache, log *logger.Logger) Cache {
	cache := Cache{
		config:    cfg,
		logger:    log,
		UserCache: &UserCache{},
	}
	return cache
}
