package storage

import (
	"SPRk_Space/internal/config"
	"SPRk_Space/internal/entities"
	"SPRk_Space/internal/logger"
	"context"
	"fmt"
	"io"
)

type StorageDriver interface {
	UserStorageDriver
	TaskStorageDriver
	io.Closer
}

type UserStorageDriver interface {
	GetUser(context.Context, int64) (entities.User, error)
	SaveUser(context.Context, entities.User) error
}

type TaskStorageDriver interface {
	GetTask(context.Context, int64) (entities.Task, error)
	SaveTask(context.Context, entities.Task) error
}

type Storage struct {
	User   UserStorageDriver
	Task   TaskStorageDriver
	closer io.Closer
}

func NewStorage(ctx context.Context, cfg *config.Storage, log logger.Logger) (*Storage, error) {
	var sd StorageDriver
	var err error
	switch cfg.StorageDriver {
	case config.DriverDB:
		sd, err = newDB(ctx, &cfg.DB, log)
	case config.DriverFakeStorage:
		sd, err = newFakeStorage(ctx, &cfg.FakeStorage, log)
	case config.DriverCache:
		sd, err = newCache(ctx, &cfg.Cache, log)
	default:
		return nil, fmt.Errorf("storage: NewStorage: unknown storage driver: %q", cfg.StorageDriver)
	}
	if err != nil {
		return nil, fmt.Errorf("storage: NewStorage: new %s driver: %w", cfg.StorageDriver, err)
	}
	storage := Storage{
		User:   sd,
		Task:   sd,
		closer: sd,
	}
	return &storage, nil
}

func (s *Storage) Close() error {
	return s.closer.Close()
}
