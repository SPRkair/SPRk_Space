package services

import (
	"SPRk_Space/internal/entities"
	"SPRk_Space/internal/logger"
	"SPRk_Space/internal/storage"
	"context"
	"math/rand/v2"
	"strconv"
)

type FakeService struct {
	ctx   context.Context
	log   logger.Logger
	Users storage.UserStorageDriver
}

func (fs *FakeService) Run() {
	fs.log.Println("Info", "FakeService Run")
	err := fs.test_get_from_cache_and_fake()
	if err != nil {
		fs.log.Printf("Error", "FakeService: Run: ", err)
	}
}

func (fs *FakeService) test_get_from_cache_and_fake() error {
	fs.log.Printf("debug_value", "test0: ", fs.Users)
	err := fs.Users.Test("1-0")
	user, err := fs.Users.GetUser(fs.ctx, 8062)
	if err != nil {
		fs.log.Printf("Error", " GetUser : %w", err)
	}
	err = fs.Users.Test("1-1")
	fs.log.Printf("debug_value", "test-2-1: ", fs.Users)
	fs.log.Printf("debug_value", "test-3-1: ", user)
	user.Name = user.Name + "!!!"
	fs.log.Printf("debug_value", "test-3-2: ", user)
	err = fs.Users.UpdateUser(fs.ctx, user)
	if err != nil {
		fs.log.Printf("Error", " UpdateUser : %w", err)
	}
	err = fs.Users.Test("1-2")
	fs.log.Printf("debug_value", "test-2-2: ", fs.Users)
	return nil
}

func (fs *FakeService) test_load_and_save_user() error {
	users, err := fs.Users.GetUsers(fs.ctx)
	if err != nil {
		fs.log.Printf("Error", "test_load_and_save_user: %v", err)
	}
	for i, user := range users {
		fs.log.Println("Value: ", i, " _ ", user)
	}
	fs.log.Println("value", "FakeService Load: ", users)
	newUsers := randUser(3)
	fs.log.Println("Info", "FakeService Rand: ", newUsers)
	users = append(users, newUsers...)
	fs.log.Println("Info", "FakeService Save")
	err = fs.Users.SaveUsers(fs.ctx, users)
	if err != nil {
		fs.log.Printf("Error", "test_load_and_save_user: %v", err)
	}
	fs.log.Println("Info", "FakeService Complete")

	return nil
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
