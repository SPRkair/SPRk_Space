package storage

import (
	"SPRk_Space/internal/config"
	"SPRk_Space/internal/entities"
	"SPRk_Space/internal/logger"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"sync"
)

type Cache struct {
	ctx    context.Context
	config *config.Cache
	logger logger.Logger

	userCache     *userCache
	taskCache     *taskCache
	loginCache    *loginCache
	sessionsCache *sessionCache
}

type userCache struct {
	mu    sync.Mutex
	users map[int64]*entities.User
}

type taskCache struct {
	mu    sync.Mutex
	tasks map[int64]*entities.Task
}

type loginCache struct {
	mu     sync.Mutex
	logins map[string]int64
}

type sessionCache struct {
	mu       sync.Mutex
	sessions map[string]*entities.Session
}

func newCache(ctx context.Context, cfg *config.Cache, log logger.Logger) (*Cache, error) {
	cache := Cache{
		ctx:           ctx,
		config:        cfg,
		logger:        log,
		userCache:     &userCache{users: make(map[int64]*entities.User)},
		taskCache:     &taskCache{tasks: make(map[int64]*entities.Task)},
		loginCache:    &loginCache{logins: make(map[string]int64)},
		sessionsCache: &sessionCache{sessions: make(map[string]*entities.Session)},
	}
	return &cache, nil
}

func (c *Cache) GetUser(ctx context.Context, id int64) (entities.User, error) {
	c.userCache.mu.Lock()
	defer c.userCache.mu.Unlock()
	u, ok := c.userCache.users[id]
	if !ok {
		return entities.User{}, ErrUserNotFound
	}
	return *u, nil
}

func (c *Cache) GetUsers(ctx context.Context) ([]entities.User, error) {
	c.userCache.mu.Lock()
	defer c.userCache.mu.Unlock()
	users := make([]entities.User, 0, len(c.userCache.users))
	for _, u := range c.userCache.users {
		users = append(users, *u)
	}
	if len(users) == 0 {
		return nil, ErrUsersNotFound
	}
	return users, nil
}

func (c *Cache) SaveUser(ctx context.Context, user entities.User) error {
	c.userCache.mu.Lock()
	defer c.userCache.mu.Unlock()
	c.userCache.users[user.Id] = &user
	return nil
}

func (c *Cache) SaveUsers(ctx context.Context, users []entities.User) error {
	c.userCache.mu.Lock()
	defer c.userCache.mu.Unlock()
	for i := range users {
		u := users[i]
		c.userCache.users[u.Id] = &u
	}
	return nil
}

func (c *Cache) UpdateUser(ctx context.Context, user entities.User) error {
	c.userCache.mu.Lock()
	defer c.userCache.mu.Unlock()
	_, ok := c.userCache.users[user.Id]
	if !ok {
		return ErrUserNotFound
	}
	c.userCache.users[user.Id] = &user
	return nil
}

func (c *Cache) DeleteUser(ctx context.Context, id int64) error {
	c.userCache.mu.Lock()
	defer c.userCache.mu.Unlock()
	delete(c.userCache.users, id)
	return nil
}

func (c *Cache) FindUserByAuthKey(key string) (int64, error) {
	c.loginCache.mu.Lock()
	defer c.loginCache.mu.Unlock()
	userId, ok := c.loginCache.logins[key]
	if !ok {
		return 0, ErrLoginNotFound
	}
	return userId, nil
}

func (c *Cache) RegNewLogin(key string) (int64, error) {
	c.loginCache.mu.Lock()
	defer c.loginCache.mu.Unlock()
	userId, ok := c.loginCache.logins[key]
	if ok {
		return 0, ErrLoginAlreadyExist
	}
	var maxID int64
	for _, id := range c.loginCache.logins {
		if id > maxID {
			maxID = id
		}
	}
	userId = maxID + 1
	c.loginCache.logins[key] = userId
	return userId, nil
}

func (c *Cache) SaveSession(ctx context.Context, session entities.Session) error {
	c.sessionsCache.mu.Lock()
	defer c.sessionsCache.mu.Unlock()
	c.sessionsCache.sessions[session.Token] = &session
	return nil
}

func (c *Cache) Verify(ctx context.Context, token string) (int64, error) {
	c.sessionsCache.mu.Lock()
	defer c.sessionsCache.mu.Unlock()
	sum := sha256.Sum256([]byte(token))
	tokenHash := hex.EncodeToString(sum[:])
	session, ok := c.sessionsCache.sessions[tokenHash]
	if !ok {
		return 0, ErrTokenNotFound
	}
	return session.UserId, nil
}

func (c *Cache) Test(string) error {
	return nil
}

func (c *Cache) GetTask(ctx context.Context, id int64) (entities.Task, error) {
	return entities.Task{}, nil
}

func (c *Cache) SaveTask(ctx context.Context, task entities.Task) error {
	return nil
}

func (c *Cache) Close() error {
	c.logger.Debug_info("Cache Close")
	return nil
}
