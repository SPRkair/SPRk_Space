package entities

import "time"

type Login struct {
	UserID int64
	Hash   string
}

type RegistrationInput struct {
	Login      string
	Pass       string
	UserName   string
	Email      string
	CreateTime time.Time
}
