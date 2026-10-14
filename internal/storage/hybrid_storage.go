package storage

import (
	"SPRk_Space/internal/config"
	"SPRk_Space/internal/entities"
	"SPRk_Space/internal/logger"
	"context"
	"fmt"
	"sync"
)

type HybridStorage struct {
	ctx    context.Context
	mu     sync.Mutex
	config *config.HybridStorage
	logger logger.Logger

	cache             *Cache
	additionalStorage *StorageDriver
}

func newHybridStorage(ctx context.Context, cfg *config.Storage, log logger.Logger) (*HybridStorage, error) {
	var additionalStorage StorageDriver
	var err error
	switch cfg.HybridStorage.AdditionalStorage {
	case config.DriverDB:
		additionalStorage, err = newDB(ctx, &cfg.DB, log)
	case config.DriverFakeStorage:
		additionalStorage, err = newFakeStorage(ctx, &cfg.FakeStorage, log)
	default:
		return nil, fmt.Errorf("cache: newCache: unknown storage driver: %q", cfg.HybridStorage.AdditionalStorage, err)
	}
	cache := Cache{
		ctx:       ctx,
		config:    &cfg.Cache,
		logger:    log,
		UserCache: &UserCache{},
	}
	hybridStorage := HybridStorage{
		ctx:    ctx,
		config: &cfg.HybridStorage,
		logger: log,

		cache:             &cache,
		additionalStorage: &additionalStorage,
	}
	return &hybridStorage, nil
}

func (hs *HybridStorage) SyncTest() {
	hs.mu.Lock()
	defer hs.mu.Unlock()
}

func (hs *HybridStorage) GetUser(ctx context.Context, id int64) (entities.User, error) {
	return entities.User{}, nil
}

func (hs *HybridStorage) GetUsers(ctx context.Context) ([]entities.User, error) {
	return nil, nil
}

func (hs *HybridStorage) SaveUser(ctx context.Context, user entities.User) error {
	return nil
}

func (hs *HybridStorage) SaveUsers(ctx context.Context, users []entities.User) error {
	return nil
}

func (hs *HybridStorage) GetTask(ctx context.Context, id int64) (entities.Task, error) {
	return entities.Task{}, nil
}

func (hs *HybridStorage) SaveTask(ctx context.Context, task entities.Task) error {
	return nil
}

func (hs *HybridStorage) Close() error {
	return nil
}
