package storage

import (
	"SPRk_Space/internal/config"
	"SPRk_Space/internal/entities"
	"SPRk_Space/internal/logger"
	"context"
	"errors"
	"fmt"
	"io"
)

var (
	ErrUserNotFound  = errors.New("user not found")
	ErrUsersNotFound = errors.New("users not found")
)

type StorageDriver interface {
	UserStorageDriver
	TaskStorageDriver
	io.Closer
}

type UserStorageDriver interface {
	GetUser(context.Context, int64) (entities.User, error)
	GetUsers(context.Context) ([]entities.User, error)
	SaveUser(context.Context, entities.User) error
	SaveUsers(context.Context, []entities.User) error
	UpdateUser(context.Context, entities.User) error
	DeleteUser(context.Context, int64) error
	Test(string) error
}

type TaskStorageDriver interface {
	GetTask(context.Context, int64) (entities.Task, error)
	SaveTask(context.Context, entities.Task) error
}

type Storage struct {
	ctx    context.Context
	config *config.Storage
	logger logger.Logger
	User   UserStorageDriver
	Task   TaskStorageDriver
	closer io.Closer
}

func NewStorage(ctx context.Context, cfg *config.Storage, log logger.Logger) (*Storage, error) {
	var sd StorageDriver
	var asd StorageDriver
	var err error
	switch cfg.StorageDriver {
	case config.DriverDB:
		log.Printf("info", "init DB...")
		sd, err = newDB(ctx, &cfg.DB, log)
	case config.DriverFakeStorage:
		log.Printf("info", "init FakeStorage...")
		sd, err = newFakeStorage(ctx, &cfg.FakeStorage, log)
	case config.DriverCache:
		log.Printf("info", "init Cache...")
		sd, err = newCache(ctx, &cfg.Cache, log)
	case config.DriverHybridStorage:
		log.Printf("info", "init hybrid Cache...")
		cache, err := newCache(ctx, &cfg.Cache, log)
		if err != nil {
			return nil, fmt.Errorf("storage: NewStorage: newCache: %q", err)
		}
		switch cfg.HybridStorage.AdditionalStorage {
		case config.DriverDB:
			log.Printf("info", "init DB...")
			asd, err = newDB(ctx, &cfg.DB, log)
		case config.DriverFakeStorage:
			log.Printf("info", "init FakeStorage...")
			asd, err = newFakeStorage(ctx, &cfg.FakeStorage, log)
		default:
			return nil, fmt.Errorf("storage: NewStorage: unknown Additional Storage driver: %q, %w", cfg.HybridStorage.AdditionalStorage, err)
		}
		sd, err = newHybridStorage(ctx, &cfg.HybridStorage, log, cache, asd)
	default:
		return nil, fmt.Errorf("storage: NewStorage: unknown storage driver: %q, %w", cfg.StorageDriver, err)
	}
	if err != nil {
		return nil, fmt.Errorf("storage: NewStorage: new %s driver: %w", cfg.StorageDriver, err)
	}
	storage := Storage{
		ctx:    ctx,
		config: cfg,
		logger: log,
		User:   sd,
		Task:   sd,
		closer: sd,
	}
	return &storage, nil
}

func (s *Storage) Close() error {
	return s.closer.Close()
}
