package storage

import (
	"SPRk_Space/internal/entities"
	"context"
	"io"
)

type StorageDriver interface {
	UserStorageDriver
	TaskStorageDriver
	LoginStorageDriver
	SessionStorageDriver
	io.Closer
}

type CacheDriver interface {
	UserStorageDriver
	TaskStorageDriver
	LoginStorageDriver
	SessionStorageDriver
	io.Closer
	SaveLogin(entities.Login) error
}

type UserStorageDriver interface {
	GetUser(context.Context, int64) (entities.User, error)
	GetUsers(context.Context) ([]entities.User, error)
	SaveUser(context.Context, entities.User) error
	SaveUsers(context.Context, []entities.User) error
	UpdateUser(context.Context, entities.User) error
	DeleteUser(context.Context, int64) error
	Test(string) error
}

type LoginStorageDriver interface {
	RegNewLogin(string) (int64, error)
	Login(string) (int64, error)
}

type TaskStorageDriver interface {
	GetTask(context.Context, int64) (entities.Task, error)
	SaveTask(context.Context, entities.Task) error
}

type SessionStorageDriver interface {
	ValidateSession(context.Context, string) (int64, error)
	SaveSession(context.Context, entities.Session) error
}
