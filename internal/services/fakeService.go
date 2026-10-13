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
	err := fs.test_load_and_save_user()
	if err != nil {
		fs.log.Printf("Error", "FakeService: Run: ", err)
	}
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
