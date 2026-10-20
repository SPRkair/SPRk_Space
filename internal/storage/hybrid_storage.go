package storage

import (
	"SPRk_Space/internal/config"
	"SPRk_Space/internal/entities"
	"SPRk_Space/internal/logger"
	"context"
	"errors"
	"fmt"
)

type HybridStorage struct {
	ctx    context.Context
	config *config.HybridStorage
	logger logger.Logger

	cache             *Cache
	additionalStorage StorageDriver
}

func newHybridStorage(ctx context.Context, cfg *config.HybridStorage, log logger.Logger, cache *Cache, additionalStorageDriver StorageDriver) (*HybridStorage, error) {
	hybridStorage := HybridStorage{
		ctx:    ctx,
		config: cfg,
		logger: log,

		cache:             cache,
		additionalStorage: additionalStorageDriver,
	}
	return &hybridStorage, nil
}

func (hs *HybridStorage) GetUser(ctx context.Context, id int64) (entities.User, error) {
	user, err := hs.cache.GetUser(ctx, id)
	if err == nil {
		return user, nil
	}
	if !errors.Is(err, ErrUserNotFound) {
		return entities.User{}, err
	}
	if errors.Is(err, ErrUserNotFound) {
		user, err = hs.additionalStorage.GetUser(ctx, id)
		if err != nil {
			return entities.User{}, err
		}
		err := hs.cache.SaveUser(ctx, user)
		if err != nil {
			hs.logger.Error("storage: HybridStorage: GetUser: %d, %v", id, err)
		}
		return user, nil
	}
	return user, nil
}

func (hs *HybridStorage) GetUsers(ctx context.Context) ([]entities.User, error) {
	users, err := hs.additionalStorage.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (hs *HybridStorage) SaveUser(ctx context.Context, user entities.User) error {
	err := hs.additionalStorage.SaveUser(ctx, user)
	if err != nil {
		return fmt.Errorf("storage: HybridStorage: additionalStorage.SaveUser: %w", err)
	}
	err = hs.cache.SaveUser(ctx, user)
	if err != nil {
		hs.logger.Error("storage: HybridStorage: cache.SaveUser: %v", err)
	}
	return nil
}

func (hs *HybridStorage) SaveUsers(ctx context.Context, users []entities.User) error {
	err := hs.additionalStorage.SaveUsers(ctx, users)
	if err != nil {
		return fmt.Errorf("storage: HybridStorage: additionalStorage.SaveUsers: %w", err)
	}
	err = hs.cache.SaveUsers(ctx, users)
	if err != nil {
		hs.logger.Error("storage: HybridStorage: cache.SaveUsers: %v", err)
	}
	return nil
}

func (hs *HybridStorage) UpdateUser(ctx context.Context, user entities.User) error {
	if err := hs.additionalStorage.UpdateUser(ctx, user); err != nil {
		return fmt.Errorf("storage: HybridStorage: additionalStorage.UpdateUser: %w", err)
	}
	if err := hs.cache.DeleteUser(ctx, user.Id); err != nil {
		hs.logger.Error("cache delete failed: id: %d, %v", user.Id, err)
	}
	return nil
}

func (hs *HybridStorage) DeleteUser(ctx context.Context, id int64) error {
	if err := hs.additionalStorage.DeleteUser(ctx, id); err != nil {
		return fmt.Errorf("storage: HybridStorage: additionalStorage.DeleteUser: %w", err)
	}
	if err := hs.cache.DeleteUser(ctx, id); err != nil {
		hs.logger.Error("cache delete failed: %d: %v", id, err)
	}
	return nil
}

func (hs *HybridStorage) Test(prefix string) error {
	hs.logger.Debug_value("Cache Test-%s", prefix)
	for i, u := range hs.cache.userCache.users {
		hs.logger.Debug_value("Cache Test-%s: value: %d : %v", prefix, i, *u)
	}
	return nil
}

func (hs *HybridStorage) GetTask(ctx context.Context, id int64) (entities.Task, error) {
	return entities.Task{}, nil
}

func (hs *HybridStorage) SaveTask(ctx context.Context, task entities.Task) error {
	return nil
}

func (hs *HybridStorage) FindUserByAuthKey(key string) (int64, error) {
	userId, err := hs.cache.FindUserByAuthKey(key)
	if err == nil {
		return userId, nil
	}
	if !errors.Is(err, ErrUserNotFound) {
		return 0, err
	}
	if errors.Is(err, ErrUserNotFound) {
		userId, err = hs.additionalStorage.FindUserByAuthKey(key)
		if err != nil {
			return 0, err
		}
		return userId, nil
	}
	return userId, nil
}

func (hs *HybridStorage) SaveSession(ctx context.Context, session entities.Session) error {
	err := hs.additionalStorage.SaveSession(ctx, session)
	if err != nil {
		return fmt.Errorf("storage: HybridStorage: additionalStorage.SaveSession: %w", err)
	}
	err = hs.cache.SaveSession(ctx, session)
	if err != nil {
		hs.logger.Error("storage: HybridStorage: cache.SaveSession: %v", err)
	}
	return nil
}

func (hs *HybridStorage) Verify(ctx context.Context, token string) (int64, error) {
	userId, err := hs.cache.Verify(ctx, token)
	if err == nil {
		return userId, nil
	}
	if !errors.Is(err, ErrTokenNotFound) {
		return 0, err
	}
	if errors.Is(err, ErrTokenNotFound) {
		userId, err = hs.additionalStorage.Verify(ctx, token)
		if err != nil {
			return 0, err
		}
		return userId, nil
	}
	return userId, nil
}

func (hs *HybridStorage) Close() error {
	hs.logger.Debug_info("HybridStorage Close")
	return nil
}
