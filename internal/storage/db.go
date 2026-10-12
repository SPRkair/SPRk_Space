package storage

import (
	"SPRk_Space/internal/config"
	"SPRk_Space/internal/entities"
	"SPRk_Space/internal/logger"
	"context"
)

type DB struct {
	ctx    context.Context
	config *config.DB
	log    logger.Logger
}

func newDB(ctx context.Context, cfg *config.DB, log logger.Logger) (*DB, error) {
	db := DB{config: cfg,
		log: log,
	}
	return &db, nil
}

func (db DB) GetUser(int64) (entities.User, error) {
	return entities.User{}, nil
}

func (dn DB) SaveUser(entities.User) error {
	return nil
}

func (db DB) GetTask(int64) (entities.Task, error) {
	return entities.Task{}, nil
}

func (dn DB) SaveTask(entities.Task) error {
	return nil
}

func (db DB) Close() error {
	return nil
}
