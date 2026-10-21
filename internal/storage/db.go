package storage

import (
	"SPRk_Space/internal/config"
	"SPRk_Space/internal/entities"
	"SPRk_Space/internal/logger"
	"context"
)

type DB struct {
	ctx    context.Context
	config *config.DB
	logger logger.Logger
}

func newDB(ctx context.Context, cfg *config.DB, log logger.Logger) (*DB, error) {
	db := DB{
		ctx:    ctx,
		config: cfg,
		logger: log,
	}
	return &db, nil
}

func (db *DB) GetUser(ctx context.Context, id int64) (entities.User, error) {
	return entities.User{}, nil
}

func (db *DB) GetUsers(ctx context.Context) ([]entities.User, error) {
	return nil, nil
}

func (dn *DB) SaveUser(ctx context.Context, user entities.User) error {
	return nil
}

func (dn *DB) SaveUsers(ctx context.Context, users []entities.User) error {
	return nil
}

func (db *DB) UpdateUser(ctx context.Context, users entities.User) error {
	return nil
}

func (db *DB) DeleteUser(ctx context.Context, id int64) error {
	return nil
}

func (db *DB) Login(key string) (int64, error) {
	return 0, nil
}

func (db *DB) RegNewLogin(key string) (int64, error) {
	return 0, nil
}

func (db *DB) SaveSession(ctx context.Context, session entities.Session) error {
	return nil
}

func (db *DB) ValidateSession(ctx context.Context, token string) (int64, error) {
	return 0, nil
}

func (db *DB) GetTask(ctx context.Context, id int64) (entities.Task, error) {
	return entities.Task{}, nil
}

func (dn *DB) SaveTask(ctx context.Context, task entities.Task) error {
	return nil
}

func (db *DB) Close() error {
	return nil
}

func (db *DB) Test(string) error {
	db.logger.Debug_info("db Close")
	return nil
}
