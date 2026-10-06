package storage

import (
	"SPRk_Space/internal/config"
	"SPRk_Space/internal/entities"
	"SPRk_Space/internal/logger"
)

type FakeStorage struct {
	config *config.FakeStorage
	logger logger.Logger

	directory string
}

func newFakeStorage(cfg *config.FakeStorage, log logger.Logger) *FakeStorage {
	fakeStorage := FakeStorage{
		config:    cfg,
		logger:    log,
		directory: cfg.DirectoryFakeStorage,
	}
	return &fakeStorage
}

func (fs *FakeStorage) GetUser(int64) (entities.User, error) {
	return entities.User{}, nil
}

func (fs *FakeStorage) SaveUser(entities.User) error {
	return nil
}
