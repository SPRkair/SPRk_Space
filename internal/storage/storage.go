package storage

import (
	"SPRk_Space/internal/config"
	"SPRk_Space/internal/entities"
	"SPRk_Space/internal/logger"
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

func (sd StorageDrivers) GetStorageDriver() StorageDriver {
	switch sd.currentDriver {
	case "DB":
		return sd.db
	case "FakeStorage":
		return sd.fakeStorage
	case "Cache":
		return sd.cache
	}
	return nil
}

type UserRepository struct {
	storageDriver StorageDriver
}

func NewStorage(cfg *config.Storage, log logger.Logger) *Storage {
	Storage := Storage{
		storageDrivers: NewStorageDrivers(cfg, log),
	}
	Storage.User = UserRepository{storageDriver: Storage.storageDrivers.GetStorageDriver()}
	return &Storage
}

func NewStorageDrivers(cfg *config.Storage, log logger.Logger) *StorageDrivers {
	cache := newCache(&cfg.Cache, log)
	db := newDB(&cfg.DB, log)
	fakestorage := newFakeStorage(&cfg.FakeStorage, log)
	StorageDrivers := StorageDrivers{
		cache:         cache,
		db:            db,
		fakeStorage:   fakestorage,
		currentDriver: cfg.StorageDriver,
	}
	return &StorageDrivers
}

func (u *UserRepository) Get(id int64) (entities.User, error) {
	User, err := u.storageDriver.GetUser(id)
	return User, err
}

func (u *UserRepository) Save(user entities.User) error {

	return u.storageDriver.SaveUser(user)
}
