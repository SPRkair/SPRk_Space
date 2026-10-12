package storage

import (
	"SPRk_Space/internal/config"
	"SPRk_Space/internal/entities"
	"SPRk_Space/internal/logger"
	"sync"
)

type Cache struct {
	mu     sync.Mutex
	config *config.Cache
	logger logger.Logger

	UserCache *UserCache
}

type UserCache struct {
	Users map[int64]*entities.User
}

func newCache(cfg *config.Cache, log logger.Logger) (*Cache, error) {
	cache := Cache{
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

func (c *Cache) GetUser(int64) (entities.User, error) {
	return entities.User{}, nil
}

func (c *Cache) SaveUser(entities.User) error {
	return nil
}

func (c *Cache) GetTask(int64) (entities.Task, error) {
	return entities.Task{}, nil
}

func (c *Cache) SaveTask(entities.Task) error {
	return nil
}

func (c *Cache) Close() error {
	return nil
}
