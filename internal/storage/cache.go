package storage

import (
	"SPRk_Space/internal/config"
	"SPRk_Space/internal/entities"
	"SPRk_Space/internal/logger"
	"context"
	"sync"
)

type Cache struct {
	ctx    context.Context
	mu     sync.Mutex
	config *config.Cache
	logger logger.Logger

	UserCache *UserCache
	TaskCache *TaskCache
}

type UserCache struct {
	Users map[int64]*entities.User
}

type TaskCache struct {
	Tasks map[int64]*entities.Task
}

func newCache(ctx context.Context, cfg *config.Cache, log logger.Logger) (*Cache, error) {
	cache := Cache{
		ctx:       ctx,
		config:    cfg,
		logger:    log,
		UserCache: &UserCache{},
	}
	return &cache, nil
}

func (c *Cache) SyncTest() {
	c.mu.Lock()
	defer c.mu.Unlock()
}

func (c *Cache) GetUser(ctx context.Context, id int64) (entities.User, error) {
	return entities.User{}, nil
}

func (c *Cache) GetUsers(ctx context.Context) ([]entities.User, error) {
	return nil, nil
}

func (c *Cache) SaveUser(ctx context.Context, user entities.User) error {
	return nil
}

func (c *Cache) SaveUsers(ctx context.Context, users []entities.User) error {
	return nil
}

func (c *Cache) GetTask(ctx context.Context, id int64) (entities.Task, error) {
	return entities.Task{}, nil
}

func (c *Cache) SaveTask(ctx context.Context, task entities.Task) error {
	return nil
}

func (c *Cache) Close() error {
	return nil
}
