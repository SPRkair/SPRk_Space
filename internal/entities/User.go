package entities

import "time"

type UserStatus string

const (
	Registered UserStatus = "Registered"
)

type User struct {
	Id           int64  `cli:"ID,ro"`
	Name         string `cli:"Name"`
	Email        string
	Status       UserStatus
	Secret_field string `cli:"-"`
	Registration time.Time
	Permissions  []Permissions
}

type Permissions struct {
	Perm     string
	Hint     string
	TypePerm string
	value    string
}
