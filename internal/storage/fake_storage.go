package storage

import (
	"SPRk_Space/internal/config"
	"SPRk_Space/internal/entities"
	"SPRk_Space/internal/logger"
	"SPRk_Space/internal/utils"
	"context"
	"fmt"
)

type FakeStorage struct {
	ctx    context.Context
	config *config.FakeStorage
	logger logger.Logger

	directory string
}

func newFakeStorage(ctx context.Context, cfg *config.FakeStorage, log logger.Logger) (*FakeStorage, error) {
	fakeStorage := FakeStorage{
		config:    cfg,
		logger:    log,
		directory: cfg.DirectoryFakeStorage,
	}
	return &fakeStorage, nil
}

func (fs *FakeStorage) GetUser(ctx context.Context, id int64) (entities.User, error) {
	users := []entities.User{}
	users, err := utils.LoadFromJSON[[]entities.User](fs.config.DirectoryFakeStorage + "users.json")
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
	users, err := utils.LoadFromJSON[[]entities.User](fs.config.DirectoryFakeStorage + "users.json")
	if err != nil {
		return nil, fmt.Errorf("FakeStorage: GetUser: %w", err)
	}
	return users, nil
}

func (fs *FakeStorage) SaveUser(ctx context.Context, user entities.User) error {
	err := utils.SaveToJSON(fs.config.DirectoryFakeStorage+"users.json", user)
	if err != nil {
		return fmt.Errorf("FakeStorage: SaveUser: %w", err)
	}
	return nil
}

func (fs *FakeStorage) SaveUsers(ctx context.Context, users []entities.User) error {
	err := utils.SaveToJSON(fs.config.DirectoryFakeStorage+"users.json", users)
	if err != nil {
		return fmt.Errorf("FakeStorage: SaveUser: %w", err)
	}
	return nil
}

func (fs *FakeStorage) GetTask(ctx context.Context, id int64) (entities.Task, error) {
	return entities.Task{}, nil
}

func (fs *FakeStorage) SaveTask(ctx context.Context, task entities.Task) error {
	return nil
}
func (fs *FakeStorage) Close() error {
	return nil
}
