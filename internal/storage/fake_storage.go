package storage

import (
	"SPRk_Space/internal/config"
	"SPRk_Space/internal/entities"
	"SPRk_Space/internal/logger"
	"SPRk_Space/internal/utils"
	"context"
	"fmt"
	"sync"
)

type FakeStorage struct {
	ctx    context.Context
	config *config.FakeStorage
	logger logger.Logger
	mu     sync.Mutex

	directory string
}

func newFakeStorage(ctx context.Context, cfg *config.FakeStorage, log logger.Logger) (*FakeStorage, error) {
	fakeStorage := FakeStorage{
		ctx:       ctx,
		config:    cfg,
		logger:    log,
		directory: cfg.DirectoryFakeStorage,
	}
	return &fakeStorage, nil
}

func (fs *FakeStorage) GetUser(ctx context.Context, id int64) (entities.User, error) {
	users := []entities.User{}
	err := utils.LoadFrom(fs.config.DirectoryFakeStorage+"users.json", &users)
	if err != nil {
		return entities.User{}, fmt.Errorf("FakeStorage: GetUser: %w", err)
	}
	for i, user := range users {
		if user.Id == id {
			return users[i], nil
		}
	}
	return entities.User{}, nil
}

func (fs *FakeStorage) GetUsers(ctx context.Context) ([]entities.User, error) {
	users := []entities.User{}
	err := utils.LoadFrom(fs.config.DirectoryFakeStorage+"users.json", &users)
	if err != nil {
		return nil, fmt.Errorf("FakeStorage: GetUser: %w", err)
	}
	return users, nil
}

func (fs *FakeStorage) SaveUser(ctx context.Context, user entities.User) error {
	err := utils.SaveTo(fs.config.DirectoryFakeStorage+"users.json", user)
	if err != nil {
		return fmt.Errorf("FakeStorage: SaveUser: %w", err)
	}
	return nil
}

func (fs *FakeStorage) SaveUsers(ctx context.Context, users []entities.User) error {
	err := utils.SaveTo(fs.config.DirectoryFakeStorage+"users.json", users)
	if err != nil {
		return fmt.Errorf("FakeStorage: SaveUser: %w", err)
	}
	return nil
}

func (fs *FakeStorage) UpdateUser(ctx context.Context, user entities.User) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	var users []entities.User
	err := utils.LoadFrom(fs.config.DirectoryFakeStorage+"users.json", &users)
	if err != nil {
		return fmt.Errorf("storage: FakeStorage: UpdateUser: LoadFrom: %d, %w", user.Id, err)
	}
	found := false
	for i, u := range users {
		if u.Id == user.Id {
			users[i] = user
			found = true
			break
		}
	}
	if !found {
		return ErrUserNotFound
	}
	err = utils.SaveTo(fs.config.DirectoryFakeStorage+"users.json", users)
	if err != nil {
		return fmt.Errorf("storage: FakeStorage: UpdateUser: SaveTo: %d, %w", user.Id, err)
	}
	return nil
}

func (fs *FakeStorage) DeleteUser(ctx context.Context, id int64) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	var users []entities.User
	err := utils.LoadFrom(fs.config.DirectoryFakeStorage+"users.json", &users)
	if err != nil {
		return fmt.Errorf("storage: FakeStorage: DeleteUser: LoadFrom: %d, %w", id, err)
	}

	target := -1
	for i := range users {
		if users[i].Id == id {
			target = i
			break
		}
	}
	if target == -1 {
		return ErrUserNotFound
	}
	users = append(users[:target], users[target+1:]...)

	err = utils.SaveTo(fs.config.DirectoryFakeStorage+"users.json", users)
	if err != nil {
		return fmt.Errorf("storage: FakeStorage: DeleteUser: SaveTo: %d, %w", id, err)
	}
	return nil
}

func (fs *FakeStorage) Test(_ string) error {
	return nil
}

func (fs *FakeStorage) GetTask(ctx context.Context, id int64) (entities.Task, error) {
	return entities.Task{}, nil
}

func (fs *FakeStorage) SaveTask(ctx context.Context, task entities.Task) error {
	return nil
}
func (fs *FakeStorage) Close() error {
	fs.logger.Debug_info("FakeStorage Close")
	return nil
}
