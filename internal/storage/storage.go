package storage

import (
	"SPRk_Space/internal/config"
	"SPRk_Space/internal/entities"
	"SPRk_Space/internal/logger"
	"fmt"
)

type StorageDriver interface {
	UserStorageDriver
	TaskStorageDriver
}

type UserStorageDriver interface {
	GetUser(int64) (entities.User, error)
	SaveUser(entities.User) error
}

type TaskStorageDriver interface {
	GetTask(int64) (entities.Task, error)
	SaveTask(entities.Task) error
}

type Storage struct {
	User *UserRepository
	Task *TaskRepository
}

type UserRepository struct {
	storageDriver UserStorageDriver
}

type TaskRepository struct {
	storageDriver TaskStorageDriver
}

func NewStorage(cfg *config.Storage, log logger.Logger) (*Storage, error) {
	var sd StorageDriver
	switch cfg.StorageDriver {
	case "DB":
		sd = newDB(&cfg.DB, log)
	case "FakeStorage":
		sd = newFakeStorage(&cfg.FakeStorage, log)
	case "Cache":
		sd = newCache(&cfg.Cache, log)
	default:
		return nil, fmt.Errorf("error: storage: GetStorageDriver: unknown storage driver: %s", cfg.StorageDriver)
	}
	storage := Storage{
		User: &UserRepository{storageDriver: sd},
		Task: &TaskRepository{storageDriver: sd},
	}
	return &storage, nil
}

func (ur *UserRepository) Get(id int64) (entities.User, error) {
	user, err := ur.storageDriver.GetUser(id)
	return user, err
}

func (ur *UserRepository) Save(user entities.User) error {
	return ur.storageDriver.SaveUser(user)
}

func (tr *TaskRepository) Get(id int64) (entities.Task, error) {
	task, err := tr.storageDriver.GetTask(id)
	return task, err
}

func (tr *TaskRepository) Save(task entities.Task) error {
	return tr.storageDriver.SaveTask(task)
}
