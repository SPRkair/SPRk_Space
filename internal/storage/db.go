package storage

import (
	"SPRk_Space/internal/config"
	"SPRk_Space/internal/entities"
	"SPRk_Space/internal/logger"
)

type DB struct {
	config *config.DB
	log    logger.Logger
}

func newDB(cfg *config.DB, log logger.Logger) *DB {
	db := DB{config: cfg,
		log: log,
	}
	return &db
}

func (db DB) GetUser(int64) (entities.User, error) {
	return entities.User{}, nil
}

func (dn DB) SaveUser(entities.User) error {
	return nil
}
