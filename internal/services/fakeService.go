package services

import (
	"SPRk_Space/internal/config"
	"SPRk_Space/internal/entities"
	"SPRk_Space/internal/logger"
	"SPRk_Space/internal/storage"
	"context"
	"fmt"
	"math/rand/v2"
	"strconv"
)

type FakeService struct {
	ctx    context.Context
	logger logger.Logger
	Users  storage.UserStorageDriver
}

func NewFakeService(ctx context.Context, cfg *config.FakeService, log logger.Logger, st storage.UserStorageDriver) *FakeService {
	return &FakeService{
		ctx:    ctx,
		logger: log,
		Users:  st,
	}
}

func (fs *FakeService) Run() {
	fs.logger.Info("FakeService Run")
	err := fs.test_get_from_cache_and_fake()
	if err != nil {
		fs.logger.Error("FakeService: Run: %v", err)
	}
}

func (fs *FakeService) ShowUsersList() ([]entities.User, error) {
	users, err := fs.Users.GetUsers(fs.ctx)
	if err != nil {
		return users, fmt.Errorf("FakeService: ShowUsersList: %v", err)
	}
	return users, nil
}

func (fs *FakeService) GetUserById(id int64) (entities.User, error) {
	user, err := fs.Users.GetUser(fs.ctx, id)
	if err != nil {
		return user, fmt.Errorf("FakeService: GetUserById: %d: %v", id, err)
	}
	return user, nil
}

func (fs *FakeService) test_get_from_cache_and_fake() error {
	fs.logger.Debug_value("test0: %v", fs.Users)
	err := fs.Users.Test("1-0")
	user, err := fs.Users.GetUser(fs.ctx, 8062)
	if err != nil {
		fs.logger.Error(" GetUser : %v", err)
	}
	err = fs.Users.Test("1-1")
	fs.logger.Debug_value("test-2-1: %v", fs.Users)
	fs.logger.Debug_value("test-3-1: %v", user)
	user.Name = user.Name + "!"
	fs.logger.Debug_value("test-3-2: %v", user)
	err = fs.Users.UpdateUser(fs.ctx, user)
	if err != nil {
		fs.logger.Error("UpdateUser : %v", err)
	}
	err = fs.Users.Test("1-2")
	fs.logger.Debug_value("test-2-2: %v", fs.Users)
	return nil
}

func (fs *FakeService) test_load_and_save_user() error {
	users, err := fs.Users.GetUsers(fs.ctx)
	if err != nil {
		fs.logger.Error("test_load_and_save_user: %v", err)
	}
	for i, user := range users {
		fs.logger.Debug_value(strconv.Itoa(i), " _ ", user)
	}
	fs.logger.Debug_value("FakeService Load: %v", users)
	newUsers := randUser(3)
	fs.logger.Info("FakeService Rand: %v", newUsers)
	users = append(users, newUsers...)
	fs.logger.Info("FakeService Save")
	err = fs.Users.SaveUsers(fs.ctx, users)
	if err != nil {
		fs.logger.Error("test_load_and_save_user: %v", err)
	}
	fs.logger.Info("FakeService Complete")

	return nil
}

func (fs *FakeService) Close() {
	fs.logger.Debug_info("FakeService Close")
}

func randUser(count int64) []entities.User {
	users := []entities.User{}
	for range count {
		num := rand.Int64N(10000) + 1
		users = append(users, entities.User{
			Id:   num,
			Name: "Test_user_" + strconv.FormatInt(num, 10),
		})
	}
	return users
}
