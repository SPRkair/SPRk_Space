package storage

import (
	"SPRk_Space/internal/config"
	"SPRk_Space/internal/entities"
	"SPRk_Space/internal/logger"
	"fmt"
)

type StorageDriver interface {
	GetUser(int64) (entities.User, error)
	SaveUser(entities.User) error
}

type Storage struct {
	User UserRepository

	storageDrivers *StorageDrivers
}

type StorageDrivers struct {
	cache         *Cache
	db            *DB
	fakeStorage   *FakeStorage
	currentDriver string
}

func (sd StorageDrivers) GetStorageDriver() (StorageDriver, error) {
	switch sd.currentDriver {
	case "DB":
		return sd.db, nil
	case "FakeStorage":
		return sd.fakeStorage, nil
	case "Cache":
		return sd.cache, nil
	default:
		return sd.cache, fmt.Errorf("error: storage: GetStorageDriver: unknown storage driver %w", sd.currentDriver)
	}
}

type UserRepository struct {
	storageDriver StorageDriver
}

func NewStorage(cfg *config.Storage, log logger.Logger) (*Storage, error) {
	storage := Storage{
		storageDrivers: NewStorageDrivers(cfg, log),
	}
	storageDriver, err := storage.storageDrivers.GetStorageDriver()
	if err != nil {
		return nil, err
	}
	storage.User = UserRepository{storageDriver: storageDriver}
	return &storage, nil
}

func NewStorageDrivers(cfg *config.Storage, log logger.Logger) *StorageDrivers {
	cache := newCache(&cfg.Cache, log)
	db := newDB(&cfg.DB, log)
	fakestorage := newFakeStorage(&cfg.FakeStorage, log)
	storageDrivers := StorageDrivers{
		cache:         cache,
		db:            db,
		fakeStorage:   fakestorage,
		currentDriver: cfg.StorageDriver,
	}
	return &storageDrivers
}

func (u *UserRepository) Get(id int64) (entities.User, error) {
	user, err := u.storageDriver.GetUser(id)
	return user, err
}

func (u *UserRepository) Save(user entities.User) error {

	return u.storageDriver.SaveUser(user)
}
