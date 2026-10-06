package storage

import (
	"SPRk_Space/internal/config"
	"SPRk_Space/internal/entities"
	"SPRk_Space/internal/logger"
)

type Cache struct {
	cfg *config.Config
	log *logger.Logger

	UserCache *UserCache
}

type UserCache struct {
	Users map[int64]*entities.User
}
