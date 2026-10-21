package entities

import "time"

type RegistrationInput struct {
	Login      string
	Pass       string
	UserName   string
	Email      string
	CreateTime time.Time
}

type LoginInput struct {
	Login string
	Pass  string
}

type ValidateSessionInput struct {
	Token       string
	Fingerprint Fingerprint
}
