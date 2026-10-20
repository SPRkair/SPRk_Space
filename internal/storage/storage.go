package storage

import (
	"SPRk_Space/internal/config"
	"SPRk_Space/internal/logger"
	"context"
	"errors"
	"fmt"
	"io"
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUsersNotFound     = errors.New("users not found")
	ErrLoginNotFound     = errors.New("Login not found")
	ErrLoginAlreadyExist = errors.New("Login already exist")
	ErrTokenNotFound     = errors.New("Token not found")
	ErrTokenCollision    = errors.New("Token collision")
	ErrInvalidInput      = errors.New("Invalid Input")
)

type Storage struct {
	ctx     context.Context
	config  *config.Storage
	logger  logger.Logger
	User    UserStorageDriver
	Task    TaskStorageDriver
	Login   LoginStorageDriver
	Session SessionStorageDriver
	closer  io.Closer
}

func NewStorage(ctx context.Context, cfg *config.Storage, log logger.Logger) (*Storage, error) {
	var sd StorageDriver
	var asd StorageDriver
	var err error
	switch cfg.StorageDriver {
	case config.DriverDB:
		log.Debug_info("init DB...")
		sd, err = newDB(ctx, &cfg.DB, log)
	case config.DriverFakeStorage:
		log.Debug_info("init FakeStorage...")
		sd, err = newFakeStorage(ctx, &cfg.FakeStorage, log)
	case config.DriverCache:
		log.Debug_info("init Cache...")
		sd, err = newCache(ctx, &cfg.Cache, log)
	case config.DriverHybridStorage:
		log.Debug_info("init hybrid Cache...")
		cache, err := newCache(ctx, &cfg.Cache, log)
		if err != nil {
			return nil, fmt.Errorf("storage: NewStorage: newCache: %q", err)
		}
		switch cfg.HybridStorage.AdditionalStorage {
		case config.DriverDB:
			log.Debug_info("init DB...")
			asd, err = newDB(ctx, &cfg.DB, log)
		case config.DriverFakeStorage:
			log.Debug_info("init FakeStorage...")
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
		ctx:     ctx,
		config:  cfg,
		logger:  log,
		User:    sd,
		Task:    sd,
		Login:   sd,
		closer:  sd,
		Session: sd,
	}
	return &storage, nil
}

func (s *Storage) Close() error {
	s.logger.Debug_info("Storage Close...")
	return s.closer.Close()
}
