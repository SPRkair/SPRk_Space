package storage

import (
	"SPRk_Space/internal/config"
	"SPRk_Space/internal/entities"
	"SPRk_Space/internal/logger"
	"context"
	"fmt"
)

type StorageDriver interface {
	UserStorageDriver
	TaskStorageDriver
}

type UserStorageDriver interface {
	GetUser(int64) (entities.User, error)
	SaveUser(entities.User) error
	Close() error
}

type TaskStorageDriver interface {
	GetTask(int64) (entities.Task, error)
	SaveTask(entities.Task) error
	Close() error
}

type Storage struct {
	User UserStorageDriver
	Task TaskStorageDriver
}

func NewStorage(ctx context.Context, cfg *config.Storage, log logger.Logger) (*Storage, error) {
	var sd StorageDriver
	var err error
	switch cfg.StorageDriver {
	case config.DriverDB:
		sd, err = newDB(ctx, &cfg.DB, log)
	case config.DriverFakeStorage:
		sd, err = newFakeStorage(&cfg.FakeStorage, log)
	case config.DriverCache:
		sd, err = newCache(&cfg.Cache, log)
	default:
		return nil, fmt.Errorf("storage: NewStorage: unknown storage driver: %s", cfg.StorageDriver)
	}
	if err != nil {
		return nil, err
	}
	storage := Storage{
		User: sd,
		Task: sd,
	}
	return &storage, nil
}

func (s *Storage) Close() []error {
	var err []error
	err = append(err, s.User.Close())
	err = append(err, s.Task.Close())
	return err
}
