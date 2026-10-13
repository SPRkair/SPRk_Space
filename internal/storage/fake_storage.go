package storage

import (
	"SPRk_Space/internal/config"
	"SPRk_Space/internal/entities"
	"SPRk_Space/internal/logger"
	"context"
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
	return entities.User{}, nil
}

func (fs *FakeStorage) SaveUser(ctx context.Context, user entities.User) error {
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
