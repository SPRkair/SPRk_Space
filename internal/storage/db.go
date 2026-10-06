package storage

import (
	"SPRk_Space/internal/config"
	"SPRk_Space/internal/logger"
)

type DB struct {
	config *config.DB
	log    logger.Logger
}

func newDB(cfg *config.DB, log logger.Logger) DB {
	db := DB{config: cfg,
		log: log,
	}
	return db
}
